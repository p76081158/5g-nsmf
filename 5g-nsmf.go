package main

import (
	"fmt"
	"os"
	// "time"
	// "strconv"

	"github.com/p76081158/5g-nsmf/modules/executor"
	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/modules/optimizer/tenantbinpack"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/scheduler"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/gnb"
	"github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	// "github.com/p76081158/5g-nsmf/api/f5gnssmf"
	// "github.com/p76081158/free5gc-nssmf"
)

// struct alias
type Packer      = slicebinpack.Packer
type Bin         = slicebinpack.Bin
type Slice       = slicebinpack.Slice
type Block       = slicebinpack.Block
type SliceDeploy = slicebinpack.SliceDeploy
type DrawBlock   = slicebinpack.DrawBlock
type UeGenerator = generator.UeGenerator

// struct var
var requestCpu       []Slice
var requestBandwidth []Slice
var ueGenerator      []UeGenerator
var bin              Bin
var access           []Slice
var reject           []Slice
var deploy_info      []SliceDeploy
var draw_info        []DrawBlock
var p                Packer

// environment var
//   1000 == 1000 milicore == 1 cpu
//   60 == 60 seconds
var ResourceLimit       = 1000
var TimeWindowSize      = 600
var TimeWindowDelay     = 100
var CPUofUserPlane      = 200
var DeployTimeBias      = 1.04
var DrawScaleRatio      = 10
var gnb_ip_dictionary   = map[string]string{"466-01-000000010": "192.168.72.51", "466-11-000000010": "192.168.72.53", "466-93-000000010": "192.168.72.55"}
var gnb_ip_B_dictionary = map[string]string{"466-01-000000010": "200", "466-11-000000010": "201", "466-93-000000010": "202"}
var RequestPattern      = "300"

// change test case of network slice request and choose algorithm of network slice bin packing
var slcieRequestCase  = "demo"
// var slcieRequestCase = "demo-1"
// var slcieRequestCase = "DataSet-4/test1"
var algorithm         = "invert-pre-order"
// var algorithm         = "pre-order"
// var algorithm         = "leaf-size"
var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)
var forecastingTime   = 1
var forecastingFinish = false
// var sort              = true
var sort              = false
var concat            = true
// var concat            = false

var cmd               = ""
var slice_request     = ""
var algo              = ""

func printInfo(osArg string) {
	fmt.Printf("Usage : %s <cmd>\n", osArg)
	fmt.Println("<cmd>\n" + 
	"get-gnb-info    : get all gNB info. in Core Network\n" +
	"restart-all-gnb : restart gNB (ueransim-gnb will not work after long time idle)\n" +
	"run-demo        : run Demo\n" +
	"run-test        : run Test Algorithm (invert-pre-order, pre-order, leaf-siz)\n" +
	"debug")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 4 {
		printInfo(os.Args[0])
		
	}
	if (os.Args[1]!="") {
		cmd = os.Args[1]
	} else {
		printInfo(os.Args[0])
	}
	if (len(os.Args) == 4) {
		if (os.Args[2]!="") {
			slice_request = os.Args[2]
		} else {
			printInfo(os.Args[0])
		}
		if (os.Args[3]!="") {
			algo = os.Args[3]
			switch algo {
			case "invert-pre-order":
			case "pre-order":
			case "leaf-size":
			default:
				fmt.Printf("run-test <Algorithm> : invert-pre-order, pre-order, leaf-size <cmd>\n")
				printInfo(os.Args[0])
			}
		} else {
			printInfo(os.Args[0])
		}
	}
	
	switch cmd {
	case "get-gnb-info":
		executor.GetgNBinfo()
	case "restart-all-gnb":
		executor.RestartAllgNB()
	case "run-demo":
		executor.RunDemo()
	case "run-test":
		executor.RunAlgorithmTest(slice_request, algo)
	case "run-tenant":
		tenantbinpack.RunTenant()
	case "debug":
		executor.RunDebug()
	default:
		printInfo(os.Args[0])
	}
}

// test example:
// ./nsrgenerator DataSet-test/test1 3 10 1000 5 0.5 10 5 0.5 600
// apply specific network slice
// f5gnssmf.ApplySliceToCoreNetwork("0x01010203", "192.168.72.51", "201", "466-01-000000010", 600, 200, 0, 60, true, 1.04)
// f5gnssmf.DeleteSliceFromCoreNetwork("0x01010203")
// nssmf.ApplyNetworkSlice("0x01010203", "192.168.72.51", "201", "466-01-000000010", 600, 200)