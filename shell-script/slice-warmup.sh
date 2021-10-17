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

# Apply Service
kubectl -n free5gc scale deployment service-$ngci-$id --replicas=1

sleep 5
# Delete Service
kubectl -n free5gc scale deployment service-$ngci-$id --replicas=0



