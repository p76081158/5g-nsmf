package executor

import (
	"fmt"
	"time"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"

	// "github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"

	// "github.com/p76081158/free5gc-nssmf"
)


func RunTenantSliceBinPack(tenant_dir string, algo string) ([]Slice, []Slice, []SliceDeploy) {
	// var slcieRequestCase  = "demo"
	var slcieRequestCase  = tenant_dir
	var algorithm         = algo
	// var algorithm         = "node-concat"
	var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)
	// gnb.GetgNBinfo()
	// fmt.Println("Restart gNB ...")
	// gnb.RestartgNB("466-01-000000010")
	// gnb.RestartgNB("466-11-000000010")
	// gnb.RestartgNB("466-93-000000010")

//	GetgNBinfo()
//	fmt.Println("Restart gNB ...")
//	RestartAllgNB()
	
	// // warm up 2 mins
	// warnup := nsrhandler.GetSliceInfo(slcieRequestCase)
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Snssai
	// 	gnb_ip      := gnb_ip_dictionary[warnup[i].Ngci]
	// 	gnb_n3_ip_B := gnb_ip_B_dictionary[warnup[i].Ngci]
	// 	ngci        := warnup[i].Ngci
	// 	cpu         := warnup[i].Cpu

	// 	// deploy to kuberenets with replicca 0
	// 	f5gnssmf.DeploySliceToCoreNetwork(slice_name, gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane)
	// 	f5gnssmf.ServiceWarmUp(slice_name, ngci)
	// }
	// fmt.Println("Warm up time ...")
	// time.Sleep(30 * time.Second)

	accept_count := 0
	reject_count := 0
	slicebinpack.Mkdir("logs/binpack/" + slcieRequestCase)

	for i := 0; i < timeWindowNumber; i++ {
		count := i + 1

		if count == forecastingTime {
			forecastingFinish = true
		}

		dt := time.Now()
		fmt.Println("")
		fmt.Println(dt.String())
		fmt.Println("")
		fmt.Println("Read Network Slice Requests ", count)
		fmt.Println("")

		// get network slice request by test case, and generate ue request pattern by network slice request
		requestCpu, requestBandwidth, ueGenerator = nsrhandler.RefreshRequestList("slice-requests/" + slcieRequestCase, "slice-forecasting/" + slcieRequestCase, count, forecastingFinish, sort)
		fmt.Println(requestCpu)
		fmt.Println(requestBandwidth)
		fmt.Println(ueGenerator)
		fmt.Println("")
		fmt.Println("Result of Network Slice Bin Packing")
		fmt.Println("")

		bin = Bin{"Resource", TimeWindowSize, ResourceLimit, requestCpu}
		p = Packer{bin, access, reject, deploy_info, draw_info}
		p.Pack(algorithm, concat, )

		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")

		accept_count += len(p.AcceptSlices)
		reject_count += len(p.RejectSlices)
//		scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
		slicebinpack.DrawBinPackResult("logs/binpack/" + slcieRequestCase, strconv.Itoa(count), p.DrawInfos, TimeWindowSize, ResourceLimit, DrawScaleRatio)

		// time.Sleep(time.Duration(TimeWindowSize + TimeWindowDelay) * time.Second)
	}
	
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Snssai
	// 	// delete from core network
	// 	f5gnssmf.DeleteSliceFromCoreNetwork(slice_name)
	// }

	fmt.Println("Accept count: ", accept_count)
	fmt.Println("Reject count: ", reject_count)
	fmt.Println("Accept rate : ", float64(accept_count) / 4000.0)
	return p.AcceptSlices, p.RejectSlices, p.DeployInfos
}

func RunTenantSliceBinPackByItem(tenant_dir string, item []Slice, algo string) ([]Slice, []Slice, []SliceDeploy) {
	// var slcieRequestCase  = "demo"
	var slcieRequestCase  = tenant_dir
	var algorithm         = algo
	// var algorithm         = "node-concat"
	// var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)
	var timeWindowNumber  = 1
	// gnb.GetgNBinfo()
	// fmt.Println("Restart gNB ...")
	// gnb.RestartgNB("466-01-000000010")
	// gnb.RestartgNB("466-11-000000010")
	// gnb.RestartgNB("466-93-000000010")

//	GetgNBinfo()
//	fmt.Println("Restart gNB ...")
//	RestartAllgNB()
	
	// // warm up 2 mins
	// warnup := nsrhandler.GetSliceInfo(slcieRequestCase)
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Snssai
	// 	gnb_ip      := gnb_ip_dictionary[warnup[i].Ngci]
	// 	gnb_n3_ip_B := gnb_ip_B_dictionary[warnup[i].Ngci]
	// 	ngci        := warnup[i].Ngci
	// 	cpu         := warnup[i].Cpu

	// 	// deploy to kuberenets with replicca 0
	// 	f5gnssmf.DeploySliceToCoreNetwork(slice_name, gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane)
	// 	f5gnssmf.ServiceWarmUp(slice_name, ngci)
	// }
	// fmt.Println("Warm up time ...")
	// time.Sleep(30 * time.Second)

	accept_count := 0
	reject_count := 0
	slicebinpack.Mkdir("logs/binpack/" + slcieRequestCase)

	for i := 0; i < timeWindowNumber; i++ {
		count := i + 1

		if count == forecastingTime {
			forecastingFinish = true
		}

		dt := time.Now()
		fmt.Println("")
		fmt.Println(dt.String())
		fmt.Println("")
		fmt.Println("Read Network Slice Requests ", count)
		fmt.Println("")

		// get network slice request by test case, and generate ue request pattern by network slice request
		// requestCpu, requestBandwidth, ueGenerator = nsrhandler.RefreshRequestList("slice-requests/" + slcieRequestCase, "slice-forecasting/" + slcieRequestCase, count, forecastingFinish, sort)
		requestCpu = item
		fmt.Println(requestCpu)
		// fmt.Println(requestBandwidth)
		// fmt.Println(ueGenerator)
		fmt.Println("")
		fmt.Println("Result of Network Slice Bin Packing")
		fmt.Println("")

		bin = Bin{"Resource", TimeWindowSize, 3000, requestCpu}
		p = Packer{bin, access, reject, deploy_info, draw_info}
		p.Pack(algorithm, concat, )

		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")

		accept_count += len(p.AcceptSlices)
		reject_count += len(p.RejectSlices)
//		scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
		slicebinpack.DrawBinPackResult("logs/binpack/" + slcieRequestCase, strconv.Itoa(count), p.DrawInfos, TimeWindowSize, 3000, DrawScaleRatio)

		// time.Sleep(time.Duration(TimeWindowSize + TimeWindowDelay) * time.Second)
	}
	
	// for i := 0; i < len(warnup); i++ {
	// 	slice_name  := warnup[i].Snssai
	// 	// delete from core network
	// 	f5gnssmf.DeleteSliceFromCoreNetwork(slice_name)
	// }

	fmt.Println("Accept count: ", accept_count)
	fmt.Println("Reject count: ", reject_count)
	fmt.Println("Accept rate : ", float64(accept_count) / 4000.0)
	return p.AcceptSlices, p.RejectSlices, p.DeployInfos
}