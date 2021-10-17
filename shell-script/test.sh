#!/bin/bash

#ngci="466-01-000000010"
#bias=$(kubectl get subnets.kubeovn.io | grep -c free5gc-n4-$ngci)
#echo $bias

ttt=$(kubectl -n free5gc get networkslices.nssmf.free5gc.com -l telecom=FET | grep -n 0x0101020f)
arr=(${ttt//:/ })
sst="010204"
tt="010204"
test=$((16#$sst - 16#$tt))
#test=$(printf "%02x\n" $sst)
echo $test
echo $ttt
echo ${arr[0]}
echo $(printf "%010d" $test)


qq=$(kubectl -n free5gc get telecoms.nso.free5gc.com | grep 466-01)
arrq=(${qq// / })
echo ${arrq[3]}



if [ -z "${arr[0]}" ]
then
    echo "Please enter cpu limit of slice service"
    exit
fi

#id="$1"
#gnb_n3_ip_b=$(( 200 + id ))
#echo $gnb_n3_ip_b
test=$(kubectl -n free5gc )
