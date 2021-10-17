#!/bin/bash

if [ -z "$1" ]
then
    echo "Please enter slice id!"
    exit
else
    slice=$1
fi

if [ -z "$2" ]
then
    sfc=false
else
    sfc=$2
fi

dir="overlays"

cd network-slice/$slice
# Delete UPF
kubectl delete -k upf-$slice/$dir/

# Delete SMF
kubectl delete -k smf-$slice/$dir/

# Delete Service
kubectl delete -k service-$slice/$dir/

# Delete UE
kubectl delete -k UERANSIM-ue-$slice/overlays/

# Delete Service Monitor
#kubectl delete -f service-monitor/

# Apply NetworkSlice Custom Resource
kubectl apply -f custom-resource/network-slice-inactive-cr.yaml

# Delete Subnet
kubectl delete -f subnet/

# Delete Network-Attachment-Definition
kubectl delete -f network-attachment-definition/