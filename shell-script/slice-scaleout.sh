#!/bin/bash

if [ -z "$1" ]
then
    echo "Please enter slice id!"
    exit
else
    id=$1
fi

if [ -z "$2" ]
then
    echo "Please enter ngci"
    exit
fi

ngci="$2"

cd network-slice/$id
# Delete UPF
kubectl -n free5gc scale deployment free5gc-upf-$id --replicas=0

# Delete SMF
kubectl -n free5gc scale deployment free5gc-smf-$id --replicas=0

# Delete Service
kubectl -n free5gc scale deployment service-$ngci-$id --replicas=0

# Delete UE
kubectl -n free5gc scale deployment free5gc-ueransim-ue-$id --replicas=0

# Delete Service Monitor
#kubectl delete -f service-monitor/

# Apply NetworkSlice Custom Resource
kubectl apply -f custom-resource/network-slice-inactive-cr.yaml
