#!/bin/bash

if [ -z "$1" ]
then
    echo "Please input sst & sd in hex format!"
    exit
fi

if [ -z "$2" ]
then
    echo "Please enter NGCI!"
    exit
fi

if [ -z "$3" ]
then
    echo "Please enter UE resource pattern!"
    exit
fi

if [ -z "$4" ]
then
    echo "Please enter UE request pattern!"
    exit
fi

id="$1"
ngci="$2"
resource_pattern="$3"
request_pattern="$4"

#
# create ue-request-generator
#

cd network-slice
mkdir -p $id
cd $id

mkdir -p UERANSIM-ue-$id/overlays

cat <<EOF > UERANSIM-ue-$id/overlays/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
bases:
- ../base
patchesStrategicMerge:
- ueransim-ue-$id-request-generator.yaml
EOF

cat <<EOF > UERANSIM-ue-$id/overlays/ueransim-ue-$id-request-generator.yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: free5gc-ueransim-ue-$id
  name: free5gc-ueransim-ue-$id
spec:
  template:
    spec:
      containers:
        - name: ue-requests-generator
          image: black842679513/ue-requests-generator:v1.0.1
          imagePullPolicy: Always
          args: ["curl service-$ngci-$id:9090", "uesimtun0", "$resource_pattern", "$request_pattern"]
          #args: ["curl cpu-test:9090", "none", "$resource_pattern", "$request_pattern"]
EOF