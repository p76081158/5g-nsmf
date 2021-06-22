package main

import (
	// "fmt"
	// "time"
	// "strconv"

	// "github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/scheduler"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/gnb"
	"github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	// "github.com/p76081158/5g-nsmf/api/f5gnssmf"
	"github.com/p76081158/free5gc-nssmf"
)

// struct alias
type Packer = slicebinpack.Packer
type Bin = slicebinpack.Bin
type Slice = slicebinpack.Slice
type Block = slicebinpack.Block
type SliceDeploy = slicebinpack.SliceDeploy
type DrawBlock = slicebinpack.DrawBlock
type UeGenerator = generator.UeGenerator

// struct var
var request []Slice
var ueGenerator []UeGenerator
var bin Bin
var access []Slice
var reject []Slice
var deploy_info []SliceDeploy
var draw_info []DrawBlock
var p Packer

// environment var
//   1000 == 1000 milicore == 1 cpu
//   60 == 60 seconds
var ResourceLimit = 1000
var TimeWindowSize = 600
var TimeWindowDelay = 100
var CPUofUserPlane = 200
var DeployTimeBias = 1.04
var DrawScaleRatio = 10
var gnb_ip_dictionary = map[string]string{"466-01-000000010": "192.168.72.51", "466-11-000000010": "192.168.72.53", "466-93-000000010": "192.168.72.55"}
var gnb_ip_B_dictionary = map[string]string{"466-01-000000010": "200", "466-11-000000010": "201", "466-93-000000010": "202"}
var RequestPattern = "300"

// change test case of network slice request and choose algorithm of network slice bin packing
// var slcieRequestCase = "demo"
var slcieRequestCase = "test-10"
// var algorithm = "right-top"
// var algorithm = "top-right"
var algorithm = "leaf-size"
//var timeWindowNumber = nsrhandler.GetTestCaseTimewindowNumber(slcieRequestCase)
var forecastingTime = 0
var forecastingFinish = false

func main() {
	// gnb.GetgNBinfo()
	// fmt.Println("Restart gNB ...")
	// gnb.RestartgNB("466-01-000000010")
	// gnb.RestartgNB("466-11-000000010")
	// gnb.RestartgNB("466-93-000000010")
	// // warm up 2 mins
	// warnup := nsrhandler.GetSliceInfo(slcieRequestCase)
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Name
	// 	gnb_ip      := gnb_ip_dictionary[warnup[i].Ngci]
	// 	gnb_n3_ip_B := gnb_ip_B_dictionary[warnup[i].Ngci]
	// 	ngci        := warnup[i].Ngci
	// 	cpu         := warnup[i].Height

	// 	// deploy to kuberenets with replicca 0
	// 	f5gnssmf.DeploySliceToCoreNetwork(slice_name, gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane)
	// 	f5gnssmf.ServiceWarmUp(slice_name, ngci)
	// }
	// fmt.Println("Warm up time ...")
	// time.Sleep(30 * time.Second)

	// accept_count := 0
	// reject_count := 0
	// slicebinpack.Mkdir(slcieRequestCase)

	// for i := 0; i < timeWindowNumber; i++ {
	// 	count := i + 1

	// 	if count == forecastingTime {
	// 		forecastingFinish = true
	// 	}

	// 	dt := time.Now()
	// 	fmt.Println("")
	// 	fmt.Println(dt.String())
	// 	fmt.Println("")
	// 	fmt.Println("Read Network Slice Requests ", count)
	// 	fmt.Println("")

	// 	// get network slice request by test case, and generate ue request pattern by network slice request
	// 	request, ueGenerator = nsrhandler.RefreshRequestList(slcieRequestCase, count, forecastingFinish)
	// 	fmt.Println(request)
	// 	fmt.Println(ueGenerator)
	// 	fmt.Println("")
	// 	fmt.Println("Result of Network Slice Bin Packing")
	// 	fmt.Println("")

	// 	bin = Bin{"Resource", TimeWindowSize, ResourceLimit, request}
	// 	p = Packer{bin, access, reject, deploy_info, draw_info}
	// 	p.Pack(algorithm)

	// 	fmt.Println("Accept Slices: ", p.AcceptSlices)
	// 	fmt.Println("Reject Slices: ", p.RejectSlices)
	// 	fmt.Println("Deploy Info:   ", p.DeployInfos)
	// 	fmt.Println("Draw Info:     ", p.DrawInfos)
	// 	fmt.Println("")

	// 	accept_count += len(p.AcceptSlices)
	// 	reject_count += len(p.RejectSlices)
	// 	// scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
	// 	slicebinpack.DrawBinPackResult(slcieRequestCase, strconv.Itoa(count), p.DrawInfos, TimeWindowSize, ResourceLimit, DrawScaleRatio)

	// 	// time.Sleep(time.Duration(TimeWindowSize + TimeWindowDelay) * time.Second)
	// }
	
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Name
	// 	// delete from core network
	// 	f5gnssmf.DeleteSliceFromCoreNetwork(slice_name)
	// }

	// fmt.Println("Accept count: ", accept_count)
	// fmt.Println("Reject count: ", reject_count)
	// fmt.Println("Accept rate : ", float64(accept_count) / 4000.0)

	// gnb.RestartgNB("466-01-000000010")
	// gnb.RestartgNB("466-11-000000010")
	// gnb.RestartgNB("466-93-000000010")
	//nssmf.ApplyNetworkSlice("0x01010203", "192.168.72.51", "200", "466-01-000000010", 600, 200)
	// nssmf.DeployNetworkSlice("0x01010203", "192.168.72.51", "200", "466-01-000000010", 600, 200)
	// nssmf.ApplyNetworkSlice("0x01010203", "466-01-000000010")
	nssmf.DeleteNetworkSlice("0x01010203")
	//f5gnssmf.ApplySliceToCoreNetwork("0x01010203", "192.168.72.51", "200", "466-01-000000010", 600, 200, 0, 60, true, 1.04)
}

// test example:

// apply specific network slice
// f5gnssmf.ApplySliceToCoreNetwork("0x01010203", "192.168.72.51", "201", "466-01-000000010", 600, 200, 0, 60, true, 1.04)
// f5gnssmf.DeleteSliceFromCoreNetwork("0x01010203")
// nssmf.ApplyNetworkSlice("0x01010203", "192.168.72.51", "201", "466-01-000000010", 600, 200)