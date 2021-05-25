package main

import (
	"fmt"
	"time"
	"strings"
	//"encoding/json"
	//"io/ioutil"

	"github.com/p76081158/5g-nsmf/slice_binpack"
	"github.com/p76081158/5g-nsmf/yamlparse"
	"github.com/p76081158/free5gc-nssmf"
	// nssmfyamlparse "github.com/p76081158/free5gc-nssmf/yamlparse"
)

// 1000 == 1000 milicore == 1 cpu
// 60 == 60 seconds
var ResourceLimit = 1000
var TimeWindowSize = 600
var CPUofUserPlane = 200
var DeployTimeBias = 1.02
var gnb_ip_dictionary = map[string]string{"466-01-000000010": "192.168.72.51", "466-11-000000010": "192.168.72.53", "466-93-000000010": "192.168.72.55"}
var gnb_ip_B_dictionary = map[string]string{"466-01-000000010": "201", "466-11-000000010": "202", "466-93-000000010": "203"}

type Packer = slice_binpack.Packer
type Bin = slice_binpack.Bin
type Slice = slice_binpack.Slice
type Block = slice_binpack.Block
type SliceDeploy = slice_binpack.SliceDeploy

// apply slice to core network (by calling the nssmf of free5gc)
func ApplySliceToCoreNetwork(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string, cpu int, core_function_cpu int, start int, duration int, end bool) {
	time.Sleep(time.Duration(float64(start) * DeployTimeBias) * time.Second)
	nssmf.ApplyNetworkSlice(snssai, gnb_ip, gnb_n3_ip_B, ngci, cpu, core_function_cpu)
	if end {
		time.Sleep(time.Duration(duration) * time.Second)
		DeleteSliceFromCoreNetwork(snssai)
	}
}

// delete slice from core network (by calling the nssmf of free5gc)
func DeleteSliceFromCoreNetwork(snssai string) {
	nssmf.DeleteNetworkSlice(snssai)
}

// modify resource and apply to core network (by calling the nssmf of free5gc)
func SliceModifyServiceCPU(snssai string, ngci string, cpu int, start int, duration int, end bool) {
	//nssmfyamlparse.ModifyCPU(snssai, "upf", cpu)
	time.Sleep(time.Duration(start) * time.Second)
	nssmf.ApplyServiceCpuChange(snssai, ngci, cpu)
	if end {
		time.Sleep(time.Duration(duration) * time.Second)
		DeleteSliceFromCoreNetwork(snssai)
	} 
}

// modify resource and apply to core network (by calling the nssmf of free5gc)
func SliceModifyBandwidth(snssai string, bandwidth int) {

}

// get slice requests in the same time_window (by reading the slice-requests/timewindow.yaml)
func RefreshRequestList(timewindowID int, forecastingFinish bool) ([]Slice){
	return yamlparse.RefreshRequestList(timewindowID, forecastingFinish)
}

// schedule network slices in the same time_window
func SlicesScheduler(slicesDeploy []SliceDeploy) {
	for i :=0; i< len(slicesDeploy); i++ {
		slice_name  := strings.Split(slicesDeploy[i].Name, "-")
		gnb_ip      := gnb_ip_dictionary[slicesDeploy[i].Ngci]
		gnb_n3_ip_B := gnb_ip_B_dictionary[slicesDeploy[i].Ngci]
		ngci        := slicesDeploy[i].Ngci
		start       := slicesDeploy[i].Start
		duration    := slicesDeploy[i].Duration
		end         := slicesDeploy[i].End
		cpu         := slicesDeploy[i].Resource

		// modify existed network slice or not
		if len(slice_name) > 1 && slice_name[1] != "1" {
			go SliceModifyServiceCPU(slice_name[0], ngci, cpu, start, duration, end)
		} else {
			go ApplySliceToCoreNetwork(slice_name[0], gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane, start, duration, end)
		}
	}
}

// example
// ApplySliceToCoreNetwork("0x01010203", "192.168.72.50", "200", "466-01-000000010", 600, CPUofUserPlane)
// DeleteSliceFromCoreNetwork("0x01010203")
// kubectl -n free5gc set resources deployment service-466-93-000000010-0x01030203 --limits cpu=200m --requests cpu=200m


func main() {
	var request []Slice
	var bin Bin
	var access []Slice
	var reject []Slice
	var deploy_info []SliceDeploy
	var p Packer

	// ApplySliceToCoreNetwork("0x01010203", "192.168.72.51", "201", "466-01-000000010", 600, CPUofUserPlane,1,1)
	// ApplySliceToCoreNetwork("0x01020203", "192.168.72.53", "202", "466-11-000000010", 600, CPUofUserPlane,1,1)
	// ApplySliceToCoreNetwork("0x01030203", "192.168.72.55", "203", "466-93-000000010", 600, CPUofUserPlane,1,1)

	// SliceModifyServiceCPU("0x01030203", "466-93-000000010", 400)

	// DeleteSliceFromCoreNetwork("0x01010203")
	// DeleteSliceFromCoreNetwork("0x01020203")
	// DeleteSliceFromCoreNetwork("0x01030203")
	// nssmfyamlparse.ModifyCPU("0x01010203", "upf", 300)

    // fmt.Println("123")

	// // stop := make(chan bool)
	// // ticker := time.NewTicker(60 * time.Second)
	timeWindowNumber := 2
	forecastingTime := 1
	forecastingFinish := false

	for i := 0; i < timeWindowNumber; i++ {
		count := i + 1
		dt := time.Now()
		fmt.Println(dt.String())
		fmt.Println("")
		fmt.Println("Read Network Slice Requests ", count)
		fmt.Println("")
		request = yamlparse.RefreshRequestList(count, forecastingFinish)
		fmt.Println(request)

		fmt.Println("")
		fmt.Println("Result of Network Slice Bin Packing")
		fmt.Println("")
		bin = Bin{"Resource", TimeWindowSize, ResourceLimit, request}
		p = Packer{bin, access, reject, deploy_info}
		p.Pack()

		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)

		SlicesScheduler(p.DeployInfos)


		if count == forecastingTime {
			forecastingFinish = true
		}
		time.Sleep(720 * time.Second)
	}
	//ApplySliceToCoreNetwork("0x01010203", "50", "200", "466-01-000000010", 600)



	//DeleteSliceFromCoreNetwork("0x01010203")
}