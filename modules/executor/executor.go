package executor

import (
	"fmt"
	"time"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/modules/optimizer/scheduler"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	"github.com/p76081158/5g-nsmf/api/f5gnssmf"
	// "github.com/p76081158/free5gc-nssmf"
)

func RunDemo() {
	var slcieRequestCase  = "demo"
	var algorithm         = "invert-pre-order"
	var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)
	gnb_ip_dictionary     = GetgNBdictionary()
	gnb_ip_B_dictionary   = GetgNB_B_dictionary()

	GetgNBinfo()
	fmt.Println("Restart gNB ...")
	RestartAllgNB()
	time.Sleep(60 * time.Second)

	// // warm up 2 mins
	warnup := nsrhandler.GetSliceInfo(slcieRequestCase)
	for i := 0; i < len(warnup); i++ {
		slice_name  := warnup[i].Snssai
		gnb_ip      := gnb_ip_dictionary[warnup[i].Ngci]
		gnb_n3_ip_B := gnb_ip_B_dictionary[warnup[i].Ngci]
		ngci        := warnup[i].Ngci
		cpu         := warnup[i].Cpu

		// deploy to kuberenets with replicca 0
		f5gnssmf.DeploySliceToCoreNetwork(slice_name, gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane)
		f5gnssmf.ServiceWarmUp(slice_name, ngci)
	}
	fmt.Println("Warm up time ...")
	time.Sleep(30 * time.Second)

	accept_count := 0
	reject_count := 0
	total_count  := 0
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

		// *** can change bin size here ***
		bin = Bin{"Resource", TimeWindowSize, ResourceLimit, requestCpu}
		p = Packer{bin, access, reject, deploy_info, draw_info}
		p.Pack(algorithm, concat)

		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")

		accept_count += len(p.AcceptSlices)
		reject_count += len(p.RejectSlices)
		total_count  += accept_count + reject_count
		
		scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
		slicebinpack.DrawBinPackResult("logs/binpack/" + slcieRequestCase, strconv.Itoa(count), p.DrawInfos, TimeWindowSize, ResourceLimit, DrawScaleRatio)

		time.Sleep(time.Duration(TimeWindowSize + TimeWindowDelay) * time.Second)
	}
	
	for i := 0; i < len(warnup); i++ {
		slice_name  := warnup[i].Snssai
		// delete from core network
		f5gnssmf.DeleteSliceFromCoreNetwork(slice_name)
	}

	fmt.Println("Accept count: ", accept_count)
	fmt.Println("Reject count: ", reject_count)
	fmt.Println("Accept rate : ", float64(accept_count) / float64(total_count))
}

func RunAlgorithmTest() {
	// var slcieRequestCase  = "demo"
	var slcieRequestCase  = "multi-tenant/466-01-000000010"
	var algorithm         = "invert-pre-order"
	// var algorithm         = "node-concat"
	var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)

	accept_count := 0
	reject_count := 0
	total_count  := 0
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
		p.Pack(algorithm, concat)

		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")

		accept_count += len(p.AcceptSlices)
		reject_count += len(p.RejectSlices)
		total_count  += accept_count + reject_count
		
		slicebinpack.DrawBinPackResult("logs/binpack/" + slcieRequestCase, strconv.Itoa(count), p.DrawInfos, TimeWindowSize, ResourceLimit, DrawScaleRatio)
	}

	fmt.Println("Accept count: ", accept_count)
	fmt.Println("Reject count: ", reject_count)
	fmt.Println("Accept rate : ", float64(accept_count) / float64(total_count))
}


// for debug
func RunDebug() {
	// f5gnssmf.ApplySliceToCoreNetwork("0x01010203", "192.168.72.51", "200", "466-01-000000010", 600, 200, 0, 60, true, 1.04)
	f5gnssmf.DeleteSliceFromCoreNetwork("0x01010203")
	f5gnssmf.DeleteSliceFromCoreNetwork("0x01020203")
	f5gnssmf.DeleteSliceFromCoreNetwork("0x01030203")
}