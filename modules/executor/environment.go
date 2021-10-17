package executor

import (
	//"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
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
// var RequestPattern      = "300"
var RequestPattern      = "1200"

// change test case of network slice request and choose algorithm of network slice bin packing
// var slcieRequestCase  = "demo"
// var slcieRequestCase = "demo-1"
// var slcieRequestCase = "DataSet-4/test1"
// var algorithm         = "invert-pre-order"
// var algorithm         = "pre-order"
// var algorithm         = "leaf-size"
// var timeWindowNumber  = nsrhandler.GetTestCaseTimewindowNumber( "slice-requests/" + slcieRequestCase)

// using forecasted data at which timewindow ( 0 == never, 1 == timewindow-1, ... )
var forecastingTime   = 2
var forecastingFinish = false
// var sort              = true
var sort              = false
var concat            = true
// var concat            = false