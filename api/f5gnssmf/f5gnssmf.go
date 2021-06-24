package f5gnssmf

import (
	"time"

	"github.com/p76081158/free5gc-nssmf"
)

// example
// ApplySliceToCoreNetwork("0x01010203", "192.168.72.50", "200", "466-01-000000010", 600, 200, 0, 60, true, 1.04)
// DeleteSliceFromCoreNetwork("0x01010203")
// kubectl -n free5gc set resources deployment service-466-93-000000010-0x01030203 --limits cpu=200m --requests cpu=200m

// deploy slice to core network (by calling the nssmf of free5gc)
func DeploySliceToCoreNetwork(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string, cpu int, core_function_cpu int) {
	nssmf.DeployNetworkSlice(snssai, gnb_ip, gnb_n3_ip_B, ngci, cpu, core_function_cpu)
}

// apply slice to core network (by calling the nssmf of free5gc)
func ApplySliceToCoreNetwork(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string, cpu int, core_function_cpu int, start int, duration int, end bool, deploy_time_bias float64) {
	time.Sleep(time.Duration(float64(start) * deploy_time_bias) * time.Second)
	nssmf.DeployNetworkSlice(snssai, gnb_ip, gnb_n3_ip_B, ngci, cpu, core_function_cpu)
	nssmf.ApplyNetworkSlice(snssai, ngci)
	if end {
		time.Sleep(time.Duration(duration) * time.Second)
		ScaleOutSliceFromCoreNetwork(snssai, ngci)
	}
}

// delete slice from core network (by calling the nssmf of free5gc)
func DeleteSliceFromCoreNetwork(snssai string) {
	nssmf.DeleteNetworkSlice(snssai)
}

// scaleout slice from core network (by calling the nssmf of free5gc)
func ScaleOutSliceFromCoreNetwork(snssai string, ngci string,) {
	nssmf.ScaleOut(snssai, ngci)
}

// warmup network slice service (by calling the nssmf of free5gc)
func ServiceWarmUp(snssai string, ngci string,) {
	nssmf.WarmUp(snssai, ngci)
}

// modify resource and apply to core network (by calling the nssmf of free5gc)
func SliceModifyServiceCPU(snssai string, ngci string, cpu int, start int, duration int, end bool) {
	time.Sleep(time.Duration(start) * time.Second)
	nssmf.ApplyServiceCpuChange(snssai, ngci, cpu)
	if end {
		time.Sleep(time.Duration(duration) * time.Second)
		ScaleOutSliceFromCoreNetwork(snssai, ngci)
	} 
}

// modify resource and apply to core network (by calling the nssmf of free5gc)
func SliceModifyBandwidth(snssai string, bandwidth int) {

}

// modify core network NFs cpu resource
// nssmfyamlparse.ModifyCPU(snssai, "upf", cpu)