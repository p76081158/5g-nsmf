#!/bin//bash

if [ -z "$1" ]
then
    echo "Please input sst & sd in hex format!"
    exit
fi

if [ -z "$2" ]
then
    echo "Please enter gnb local ip"
    exit
fi

if [ -z "$3" ]
then
    echo "Please enter gnb n3 ip"
    exit
fi

if [ -z "$4" ]
then
    echo "Please enter ngci"
    exit
fi

if [ -z "$5" ]
then
    echo "Please enter cpu limit of slice service"
    exit
fi

if [ -z "$6" ]
then
    echo "Please enter cpu limit of core network user plane network function"
    exit
fi

#if [ -d network-slice/$1 ]
#then
#    echo "Network Slice already exists"
#    exit
#fi

nsi="1"
sst=${1:2:2}
sd=${1:4}
id="$1"
gnb_ip="$2"
gnb_n3_ip="$3"
ngci="$4"
cpu="$5"
cpu_limit="${cpu}m"
core_network_function_cpu="$6"
core_network_function_cpu_limit="${core_network_function_cpu}m"
mcc=${4:0:3}
mnc=${4:4:2}

ue_cmd=$(kubectl -n free5gc get networkslices.nssmf.free5gc.com -l mcc="$mcc" -l mnc="$mnc" | grep -n $1)
ue_arr=(${ue_cmd//:/ })

if [ -z "${ue_arr[0]}" ]
then
    echo "Please check slice id"
    exit
fi

ue_id=$(( ue_arr[0] - 1 ))
ue_imsi=$(printf "%010d" $ue_id)

telecom_cmd=$(kubectl -n free5gc get telecoms.nso.free5gc.com | grep $mcc-$mnc)
telecom_arr=(${telecom_cmd// / })
abb=${telecom_arr[3]}

# bias=$(kubectl get subnets.kubeovn.io | grep -c free5gc-n4-$ngci)
ue_ip=$(( 59 + ue_id ))
n3_ip=$(( 3 + ue_id ))
n4_ip=$(( 100 + ue_id ))

dir="overlays"
start=`date +%s`
echo "Create Slice"
echo "mcc:"$mcc
echo "mnc:"$mnc
echo "nsi:"$nsi
echo "sst:"$sst
echo "sd :"$sd
echo "ue :"${mcc}${mnc}${ue_imsi}

cd network-slice
mkdir -p $id
cd $id

#
# create custom resource yaml
#

mkdir -p custom-resource

cat <<EOF > custom-resource/network-slice-active-cr.yaml
---
apiVersion: "nssmf.free5gc.com/v1"
kind: NetworkSlice
metadata:
  name: "$id"
  namespace: free5gc
  labels:
    nsi: "$nsi"
    mcc: "$mcc"
    mnc: "$mnc"
    telecom: $abb
spec:
  sst: "$sst"
  sd: "$sd"
  status: Active
  n4_cidr: "10.$gnb_n3_ip.$n4_ip.0/24"
  ue_subnet: "60.$ue_ip.0.0/16"
  cpu: $cpu_limit
  memory: Default
  bandwidth: Default
EOF

cat <<EOF > custom-resource/network-slice-inactive-cr.yaml
---
apiVersion: "nssmf.free5gc.com/v1"
kind: NetworkSlice
metadata:
  name: "$id"
  namespace: free5gc
  labels:
    nsi: "$nsi"
    mcc: "$mcc"
    mnc: "$mnc"
    telecom: $abb
spec:
  sst: "$sst"
  sd: "$sd"
  status: Inactive
  n4_cidr: Undefined
  ue_subnet: Undefined
  cpu: Default
  memory: Default
  bandwidth: Default
EOF

#
# create smf yaml
#

mkdir -p smf-$id
mkdir -p smf-$id/base
mkdir -p smf-$id/base/config
mkdir -p smf-$id/overlays
cp -r ../../TLS smf-$id/base

cat <<EOF > smf-$id/base/config/smfcfg-$id.yaml
info:
  version: 1.0.0
  description: AMF initial local configuration

configuration:
  smfName: SMF
  sbi:
    scheme: http
    registerIPv4: free5gc-smf-$id # IP used to register to NRF
    bindingIPv4: 0.0.0.0  # IP used to bind the service
    port: 8000
    tls:
      key: free5gc/support/TLS/smf.key
      pem: free5gc/support/TLS/smf.pem
  serviceNameList:
    - nsmf-pdusession
    - nsmf-event-exposure
    - nsmf-oam
  snssaiInfos:
    - sNssai:
        sst: $((16#$sst))
        sd: $sd
      dnnInfos:
        - dnn: internet
          dns:
            ipv4: 8.8.8.8
            ipv6: 2001:4860:4860::8888
          ueSubnet: 60.$ue_ip.0.0/16
  pfcp:
    addr: 10.$gnb_n3_ip.$n4_ip.20
  userplane_information:
    up_nodes:
      gNB1:
        type: AN
        an_ip: $gnb_ip
      AnchorUPF1:
        type: UPF
        node_id: 10.$gnb_n3_ip.$n4_ip.101 # the IP/FQDN of N4 interface on this UPF (PFCP)
        sNssaiUpfInfos:
          - sNssai:
              sst: $((16#$sst))
              sd: $sd
            dnnUpfInfoList:
              - dnn: internet
        interfaces:
          - interfaceType: N3
            endpoints: # the IP address of this N3/N9 interface on this UPF
              - 10.$gnb_n3_ip.100.$n3_ip
            networkInstance: internet
          - interfaceType: N9
            endpoints: # the IP address of this N3/N9 interface on this UPF
              - 10.$gnb_n3_ip.$n4_ip.101
            networkInstance: internet
    links:
      - A: gNB1
        B: AnchorUPF1
  dnn:
    internet:
      dns:
        ipv4: 8.8.8.8
        ipv6: 2001:4860:4860::8888
  ue_subnet: 60.$ue_ip.0.0/16
  nrfUri: http://free5gc-nrf-$mcc-$mnc:8000
  ulcl: false

logger:
  SMF:
    debugLevel: info
    ReportCaller: false
  NAS:
    debugLevel: info
    ReportCaller: false
  NGAP:
    debugLevel: info
    ReportCaller: false
  Aper:
    debugLevel: info
    ReportCaller: false
  PathUtil:
    debugLevel: info
    ReportCaller: false
  OpenApi:
    debugLevel: info
    ReportCaller: false
  PFCP:
    debugLevel: info
    ReportCaller: false
EOF

cat <<EOF > smf-$id/base/config/uerouting.yaml
info:
  version: 1.0.0
  description: Routing information for UE

ueRoutingInfo: # the list of UE routing information
  - SUPI: imsi-${mcc}${mnc}00007487 # Subscription Permanent Identifier of the UE
    AN: $gnb_ip # the IP address of RAN (gNB)
    PathList: # the pre-config paths for this SUPI
      - DestinationIP: 60.60.0.100 # the destination IP address on Data Network (DN)
        # the order of UPF nodes in this path. We use the UPF's name to represent each UPF node.
        # The UPF's name should be consistent with smfcfg.yaml
        UPF: !!seq
          - BranchingUPF
          - AnchorUPF1

      - DestinationIP: 60.60.0.101 # the destination IP address on Data Network (DN)
        # the order of UPF nodes in this path. We use the UPF's name to represent each UPF node.
        # The UPF's name should be consistent with smfcfg.yaml
        UPF: !!seq
          - BranchingUPF
          - AnchorUPF2

  - SUPI: imsi-${mcc}${mnc}00007486 # Subscription Permanent Identifier of the UE
    AN: $gnb_ip # the IP address of RAN
    PathList: # the pre-config paths for this SUPI
      - DestinationIP: 10.10.0.10 # the destination IP address on Data Network (DN)
        # the order of UPF nodes in this path. We use the UPF's name to represent each UPF node.
        # The UPF's name should be consistent with smfcfg.yaml
        UPF: !!seq
          - BranchingUPF
          - AnchorUPF1

      - DestinationIP: 10.10.0.11 # the destination IP address on Data Network (DN)
        # the order of UPF nodes in this path. We use the UPF's name to represent each UPF node.
        # The UPF's name should be consistent with smfcfg.yaml
        UPF: !!seq
          - BranchingUPF
          - AnchorUPF2
EOF

cat <<EOF > smf-$id/base/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: free5gc
resources:
  - smf-$id-sa.yaml
  - smf-$id-rbac.yaml
  - smf-$id-service-sbi.yaml
  - smf-$id-service-n4-endpoint.yaml
  - smf-$id-service-n4.yaml
  - smf-$id-deployment.yaml

# declare Secret from a secretGenerator
secretGenerator:
- name: free5gc-smf-$id-tls-secret
  namespace: free5gc
  files:
  - TLS/smf.pem
  - TLS/smf.key
  type: "Opaque"
generatorOptions:
  disableNameSuffixHash: true

# declare ConfigMap from a ConfigMapGenerator
configMapGenerator:
- name: free5gc-smf-$id-config
  namespace: free5gc
  files:
    - smfcfg.yaml=config/smfcfg-$id.yaml
    - config/uerouting.yaml
EOF

cat <<EOF > smf-$id/base/smf-$id-sa.yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: free5gc-smf-$id-sa
EOF

cat <<EOF > smf-$id/base/smf-$id-rbac.yaml
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: free5gc-smf-$id-rbac
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: free5gc-smf-$id-sa
EOF

cat <<EOF > smf-$id/base/smf-$id-service-sbi.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: free5gc-smf-$id
  name: free5gc-smf-$id
spec:
  type: ClusterIP
  ports:
  - name: free5gc-smf-$id-sbi
    port: 8000
    protocol: TCP
    targetPort: 8000
  - name: free5gc-smf-$id-n4
    port: 8805
    protocol: UDP
    targetPort: 8805
  selector:
    app: free5gc-smf-$id
EOF

cat <<EOF > smf-$id/base/smf-$id-service-n4-endpoint.yaml
---
kind: Endpoints
apiVersion: v1
metadata:
  name: free5gc-smf-$id-n4
subsets:
  - addresses:
      - ip: 10.$gnb_n3_ip.$n4_ip.20
    ports:
      - port: 8805
EOF

cat <<EOF > smf-$id/base/smf-$id-service-n4.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: free5gc-smf-$id-n4
  name: free5gc-smf-$id-n4
spec:
  # type: ClusterIP
  # type: NodePort
  clusterIP: None
  ports:
  - name: free5gc-smf-$id-n4
    port: 8805
    protocol: UDP
    targetPort: 8805
EOF

cat <<EOF > smf-$id/base/smf-$id-deployment.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: free5gc-smf-$id
  labels:
    app: free5gc-smf-$id
    nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
    sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
    sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
spec:
  replicas: 0
  selector:
    matchLabels:
      app: free5gc-smf-$id
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: free5gc-smf-$id
        nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
        sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
        sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
      annotations:
        k8s.v1.cni.cncf.io/networks: free5gc-n4-$ngci-$id
        free5gc-n4-$ngci-$id.free5gc.ovn.kubernetes.io/logical_switch: free5gc-n4-$ngci-$id
        free5gc-n4-$ngci-$id.free5gc.ovn.kubernetes.io/ip_address: 10.$gnb_n3_ip.$n4_ip.20
    spec:
      securityContext:
        runAsUser: 0
        runAsGroup: 0
      containers:
        - name: free5gc-smf
          image: black842679513/free5gc-smf:v3.0.5
          imagePullPolicy: IfNotPresent
          # imagePullPolicy: Always
          securityContext:
            privileged: false
          volumeMounts:
            - name: free5gc-smf-$id-config
              mountPath: /free5gc/config
            - name: free5gc-smf-$id-cert
              mountPath: /free5gc/support/TLS
          ports:
            - containerPort: 8000
              name: if-sbi
              protocol: TCP
            - containerPort: 8805
              name: if-n4
              protocol: UDP
        - name: tcpdump
          image: corfr/tcpdump
          imagePullPolicy: IfNotPresent
          command:
            - /bin/sleep
            - infinity
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      serviceAccountName: free5gc-smf-$id-sa
      terminationGracePeriodSeconds: 30
      volumes:
        - name: free5gc-smf-$id-cert
          secret:
            secretName: free5gc-smf-$id-tls-secret
        - name: free5gc-smf-$id-config
          configMap:
            name: free5gc-smf-$id-config
EOF

cat <<EOF > smf-$id/overlays/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
bases:
- ../base
patchesStrategicMerge:
- smf-$id-cpu.yaml
EOF

cat <<EOF > smf-$id/overlays/smf-$id-cpu.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: free5gc-smf-$id
  name: free5gc-smf-$id
spec:
  template:
    spec:
      containers:
        - name: free5gc-smf
          resources:
            requests:
              cpu: "$core_network_function_cpu_limit"
            limits:
              cpu: "$core_network_function_cpu_limit"
EOF

#
# create upf yaml
#

mkdir -p upf-$id
mkdir -p upf-$id/base
mkdir -p upf-$id/base/config
mkdir -p upf-$id/overlays

cat <<EOF > upf-$id/base/config/upfcfg-$id.yaml
info:
  version: 1.0.0
  description: UPF configuration

configuration:
  # debugLevel: panic|fatal|error|warn|info|debug|trace
  debugLevel: info

  pfcp:
    - addr: 10.$gnb_n3_ip.$n4_ip.101

  gtpu:
    - addr: 10.$gnb_n3_ip.100.$n3_ip
    # [optional] gtpu.name
    # - name: upf.5gc.nctu.me
    # [optional] gtpu.ifname
    # - ifname: gtpif

  dnn_list:
    - dnn: internet
      cidr: 60.$ue_ip.0.0/16
      # [optional] apn_list[*].natifname
      # natifname: eth0
EOF

cat <<EOF > upf-$id/base/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: free5gc
resources:
  - upf-$id-sa.yaml
  - upf-$id-rbac.yaml
  - upf-$id-service.yaml
  - upf-$id-deployment.yaml

# declare ConfigMap from a ConfigMapGenerator
configMapGenerator:
- name: free5gc-upf-$id-config
  namespace: free5gc
  files:
    - upfcfg.yaml=config/upfcfg-$id.yaml
EOF

cat <<EOF > upf-$id/base/upf-$id-sa.yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: free5gc-upf-$id-sa
EOF

cat <<EOF > upf-$id/base/upf-$id-rbac.yaml
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: free5gc-upf-$id-rbac
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: free5gc-upf-$id-sa
EOF

cat <<EOF > upf-$id/base/upf-$id-service.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: free5gc-upf-$id
  name: free5gc-upf-$id
spec:
  type: ClusterIP
  ports:
  - name: free5gc-upf-$id-n3
    port: 2152
    protocol: UDP
    targetPort: 2152
  - name: free5gc-upf-$id-n4
    port:  8805
    protocol: UDP
    targetPort: 8805
  selector:
    app: free5gc-upf-$id
EOF

cat <<EOF > upf-$id/base/upf-$id-deployment.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: free5gc-upf-$id
  labels:
    app: free5gc-upf-$id
    nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
    sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
    sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
spec:
  replicas: 0
  selector:
    matchLabels:
      app: free5gc-upf-$id
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: free5gc-upf-$id
        nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
        sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
        sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
      annotations:
        k8s.v1.cni.cncf.io/networks: free5gc-n3-$ngci, free5gc-n4-$ngci-$id
        free5gc-n3-$ngci.free5gc.ovn.kubernetes.io/logical_switch: free5gc-n3-$ngci
        free5gc-n3-$ngci.free5gc.ovn.kubernetes.io/ip_address: 10.$gnb_n3_ip.100.$n3_ip
        free5gc-n4-$ngci-$id.free5gc.ovn.kubernetes.io/logical_switch: free5gc-n4-$ngci-$id
        free5gc-n4-$ngci-$id.free5gc.ovn.kubernetes.io/ip_address: 10.$gnb_n3_ip.$n4_ip.101
    spec:
      securityContext:
        runAsUser: 0
        runAsGroup: 0
      initContainers:
        - name: free5gc-upf-init
          image: black842679513/free5gc-upf-init:v1.0.0
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: false
            # add network capabilities
            capabilities:
              add: ["NET_ADMIN", "NET_RAW", "NET_BIND_SERVICE", "SYS_TIME"]
          command:
            - /bin/sh
            - -c
            - |
              sysctl -w net.ipv4.ip_forward=1
              iptables -t nat -A POSTROUTING -s 60.$ue_ip.0.0/16 ! -o upfgtp -j MASQUERADE
      containers:
        - name: free5gc-upf
          image: black842679513/free5gc-upf:v3.0.5
          imagePullPolicy: IfNotPresent
          # imagePullPolicy: Always
          securityContext:
            privileged: false
            # add network capabilities
            capabilities:
              add: ["NET_ADMIN", "NET_RAW", "NET_BIND_SERVICE", "SYS_TIME"]
          volumeMounts:
            - name: free5gc-upf-$id-config
              mountPath: /free5gc/config
              # read host linux tun/tap packets
           # - name: tun-dev-dir
           #   mountPath: /dev/net/tun
          ports:
            - containerPort: 2152
              name: if-n3
              protocol: UDP
            - containerPort: 8805
              name: if-n4
              protocol: UDP
        - name: tcpdump
          image: corfr/tcpdump
          securityContext:
            privileged: true
          command:
            - /bin/sh
            - -c
            - |
              sleep infinity
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      serviceAccountName: free5gc-upf-$id-sa
      terminationGracePeriodSeconds: 30
      volumes:
        - name: free5gc-upf-$id-config
          configMap:
            name: free5gc-upf-$id-config
       # - name: tun-dev-dir
       #   hostPath:
       #     path: /dev/net/tun
EOF

cat <<EOF > upf-$id/overlays/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
bases:
- ../base
patchesStrategicMerge:
- upf-$id-cpu.yaml
EOF

cat <<EOF > upf-$id/overlays/upf-$id-cpu.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: free5gc-upf-$id
  name: free5gc-upf-$id
spec:
  template:
    spec:
      containers:
        - name: free5gc-upf
          resources:
            requests:
              cpu: "$core_network_function_cpu_limit"
            limits:
              cpu: "$core_network_function_cpu_limit"
EOF

#
# create service yaml
#

mkdir -p service-$id/base
mkdir -p service-$id/overlays

cat <<EOF > service-$id/base/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: free5gc
resources:
  - service-$id-service.yaml
  - service-$id-deployment.yaml
EOF

cat <<EOF > service-$id/base/service-$id-service.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: service-$ngci-$id
  name: service-$ngci-$id
spec:
  type: ClusterIP
  ports:
  - name: service-port
    port: 9090
    protocol: TCP
    targetPort: 9090
  selector:
    app: service-$ngci-$id
EOF

cat <<EOF > service-$id/base/service-$id-deployment.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: service-$ngci-$id
  labels:
    app: service-$ngci-$id
spec:
  replicas: 0
  selector:
    matchLabels:
      app: service-$ngci-$id
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
  # minReadySeconds: 1
  template:
    metadata:
      labels:
        app: service-$ngci-$id
        nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
        sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
        sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
      annotations:
        # ovn.kubernetes.io/ingress_rate: "3"
        # ovn.kubernetes.io/egress_rate: "1"
    spec:
      containers:
        - name: cpu-usage-simulator
          image: black842679513/cpu-usage-simulator:v1.0.0
          imagePullPolicy: IfNotPresent
          # imagePullPolicy: Always
          args: ["--", "1", "1", "1"]    # ["--", stress_cpu_nums, cpu_loading, time_duration]
          securityContext:
            privileged: true
          ports:
            - containerPort: 9090
              name: service-port
              protocol: TCP
EOF

cat <<EOF > service-$id/overlays/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
bases:
- ../base
patchesStrategicMerge:
- service-$id-cpu.yaml
EOF

cat <<EOF > service-$id/overlays/service-$id-cpu.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: service-$ngci-$id
  name: service-$ngci-$id
spec:
  template:
    spec:
      containers:
        - name: cpu-usage-simulator
          resources:
            requests:
              cpu: "$cpu_limit"
            limits:
              cpu: "$cpu_limit"
EOF

#
# create ue for this slice
#

mkdir -p UERANSIM-ue-$id/base/config

cat <<EOF > UERANSIM-ue-$id/base/config/free5gc-ue.yaml
# IMSI number of the UE. IMSI = [MCC|MNC|MSISDN] (In total 15 or 16 digits)
supi: 'imsi-${mcc}${mnc}${ue_imsi}'
# Mobile Country Code value of HPLMN
mcc: '$mcc'
# Mobile Network Code value of HPLMN (2 or 3 digits)
mnc: '$mnc'

# Permanent subscription key
key: '8baf473f2f8fd09487cccbd7097c6862'
# Operator code (OP or OPC) of the UE
op: '8e27b6af0e692e750f32667a3b14605d'
# This value specifies the OP type and it can be either 'OP' or 'OPC'
opType: 'OPC'
# Authentication Management Field (AMF) value
amf: '8000'
# IMEI number of the device. It is used if no SUPI is provided
imei: '356938035643803'
# IMEISV number of the device. It is used if no SUPI and IMEI is provided
imeiSv: '4370816125816151'

# List of gNB IP addresses for Radio Link Simulation
gnbSearchList:
  - $gnb_ip

# UAC Access Identities Configuration
uacAic:
  mps: false
  mcs: false

# UAC Access Control Class
uacAcc:
  normalClass: 0
  class11: false
  class12: false
  class13: false
  class14: false
  class15: false

# Initial PDU sessions to be established
sessions:
  - type: 'IPv4'
    apn: 'internet'
    slice:
      sst: 0x$sst
      sd: 0x$sd

# Configured NSSAI for this UE by HPLMN
configured-nssai:
  - sst: 0x$sst
    sd: 0x$sd

# Default Configured NSSAI for this UE
default-nssai:
  - sst: 1
    sd: 1

# Supported encryption algorithms by this UE
integrity:
  IA1: true
  IA2: true
  IA3: true

# Supported integrity algorithms by this UE
ciphering:
  EA1: true
  EA2: true
  EA3: true

# Integrity protection maximum data rate for user plane
integrityMaxRate:
  uplink: 'full'
  downlink: 'full'
EOF

cat <<EOF > UERANSIM-ue-$id/base/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: free5gc
resources:
  - ueransim-ue-sa.yaml
  - ueransim-ue-rbac.yaml
  # - ueransim-ue-service.yaml
  - ueransim-ue-deployment.yaml

# declare ConfigMap from a ConfigMapGenerator
configMapGenerator:
- name: free5gc-ueransim-ue-$id-config
  namespace: free5gc
  files:
    - config/free5gc-ue.yaml
EOF

cat <<EOF > UERANSIM-ue-$id/base/ueransim-ue-sa.yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: free5gc-ueransim-ue-$id-sa
EOF

cat <<EOF > UERANSIM-ue-$id/base/ueransim-ue-rbac.yaml
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: free5gc-ueransim-ue-$id-rbac
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: free5gc-ueransim-ue-$id-sa
EOF

#cat <<EOF > UERANSIM-ue-$id/base/ueransim-ue-service.yaml
#EOF

cat <<EOF > UERANSIM-ue-$id/base/ueransim-ue-deployment.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: free5gc-ueransim-ue-$id
  labels:
    app: free5gc-ueransim-ue-$id
spec:
  replicas: 0
  selector:
    matchLabels:
      app: free5gc-ueransim-ue-$id
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: free5gc-ueransim-ue-$id
        mcc: "$mcc"
        mnc: "$mnc"
      annotations:
        k8s.v1.cni.cncf.io/networks: free5gc-macvlan
        # free5gc-macvlan.free5gc.kubernetes.io/ip_address: 192.168.72.60
    spec:
      securityContext:
        runAsUser: 0
        runAsGroup: 0
      containers:
        - name: free5gc-ueransim-ue
          image: black842679513/free5gc-ueransim:v3.2.0
          imagePullPolicy: IfNotPresent
          # imagePullPolicy: Always
          command:
            - /bin/bash
            - -c
            - build/nr-ue -c config/free5gc-ue.yaml
          tty: true
          securityContext:
            # allow container to access the host's resources
            privileged: true
            capabilities:
              add: ["NET_ADMIN", "SYS_TIME"]
          volumeMounts:
            - name: free5gc-ueransim-ue-$id-config
              mountPath: /UERANSIM/config
              # read host linux tun/tap packets
            #- name: tun-dev-dir  
            #  mountPath: /dev/net/tun
        - name: tcpdump
          image: corfr/tcpdump
          command:
            - /bin/sleep
            - infinity
        - name: ue-requests-generator
          image: black842679513/ue-requests-generator:v1.0.1
          imagePullPolicy: IfNotPresent
          # imagePullPolicy: Always
          args: ["curl service-$ngci-$id:9090", "uesimtun0", "500:10,400:15", 500]
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      serviceAccountName: free5gc-ueransim-ue-$id-sa
      terminationGracePeriodSeconds: 30
      volumes:
        - name: free5gc-ueransim-ue-$id-config
          configMap:
            name: free5gc-ueransim-ue-$id-config
        #- name: tun-dev-dir
        #  hostPath:
        #    path: /dev/net/tun
EOF

#
# create n4 subnet yaml
#

mkdir -p subnet

cat <<EOF > subnet/free5gc-n4.yaml
---
apiVersion: kubeovn.io/v1
kind: Subnet
metadata:
  name: free5gc-n4-$ngci-$id
  namespace: free5gc
  labels:
    nsi: "$nsi"        # Network Slice Instance of three networks (RAN,TN,CN)
    sst: "$sst"       # Slice/Service Type (1 byte uinteger, range: 0~255)
    sd: "$sd"    # Slice Differentiator (3 bytes hex string, range: 000000~FFFFFF)
spec:
  protocol: IPv4
  cidrBlock: 10.$gnb_n3_ip.$n4_ip.0/24
  gateway: 10.$gnb_n3_ip.$n4_ip.1
  excludeIps:
  - 10.$gnb_n3_ip.$n4_ip.0..10.$gnb_n3_ip.$n4_ip.10
EOF

#
# create n4 network-attachment-definition yaml
#

mkdir -p network-attachment-definition

cat <<EOF > network-attachment-definition/free5gc-n4.yaml
---
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: free5gc-n4-$ngci-$id
  namespace: free5gc
spec:
  config: '{
      "cniVersion": "0.3.1",
      "type": "kube-ovn",
      "server_socket": "/run/openvswitch/kube-ovn-daemon.sock",
      "provider": "free5gc-n4-$ngci-$id.free5gc.ovn"
    }'
EOF

#
# create Service Monitor
#

#mkdir -p service-monitor

#cat <<EOF > service-monitor/service-$id-service-monitor.yaml
#---
#apiVersion: monitoring.coreos.com/v1
#kind: ServiceMonitor
#metadata:
#  name: service-$ngci-$id-service-monitor
  # Change this to the namespace the Prometheus instance is running in
#  namespace: monitoring
#  labels:
#    app: service-$ngci-$id
#    release: prometheus
#spec:
# selector:
#    matchLabels:
#      app: service-$ngci-$id # target service
#  endpoints:
#  - port: metrics
#    interval: 5s
#EOF

#
# deploy to kubernetes
#

# c NetworkSlice Custom Resource
kubectl apply -f custom-resource/network-slice-active-cr.yaml

# Deploy Subnet
kubectl apply -f subnet/

# Deploy Network-Attachment-Definition
kubectl apply -f network-attachment-definition/

# Deploy UPF
kubectl apply -k upf-$id/$dir/

# Deploy SMF
kubectl apply -k smf-$id/$dir/

# Deploy Service
kubectl apply -k service-$id/$dir/

# Deploy UE
kubectl apply -k UERANSIM-ue-$id/overlays/

# Deploy Service Monitor
#kubectl apply -f service-monitor/

# Time of Network Slice ready
end=`date +%s`
runtime=$(( end - start ))
echo "Deployment time of Network Slice:"$runtime