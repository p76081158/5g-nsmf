#!/bin//bash

if [ -z "$1" ]
then
    echo "Please input sst & sd in hex format!"
    exit
fi

if [ -z "$2" ]
then
    echo "Please enter ngci"
    exit
fi

id="$1"
ngci="$2"
start=`date +%s`

cd network-slice/$id

# Apply NetworkSlice Custom Resource
kubectl apply -f custom-resource/network-slice-active-cr.yaml

# Apply UPF
kubectl -n free5gc scale deployment free5gc-upf-$id --replicas=1
sleep 1 # wait 1s for pod creation start, so that kubectl wait will not get error
kubectl -n free5gc wait --for=condition=ready pod -l app=free5gc-upf-$id

# Apply SMF
kubectl -n free5gc scale deployment free5gc-smf-$id --replicas=1
sleep 1
kubectl -n free5gc wait --for=condition=ready pod -l app=free5gc-smf-$id

# Apply Service
kubectl -n free5gc scale deployment service-$ngci-$id --replicas=1
sleep 1
kubectl -n free5gc wait --for=condition=ready pod -l app=service-$ngci-$id

sleep 1
# Apply UE
kubectl -n free5gc scale deployment free5gc-ueransim-ue-$id --replicas=1

# Apply Service Monitor
#kubectl apply -f service-monitor/

# Time of Network Slice ready
end=`date +%s`
runtime=$(( end - start ))
echo "Ready time of Network Slice:"$runtime