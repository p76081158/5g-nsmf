package main

import (
	"fmt"
	"os"
	// "time"
	// "strconv"

	// "github.com/p76081158/5g-nsmf/modules/nsrhandler"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/scheduler"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/gnb"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	// "github.com/p76081158/5g-nsmf/api/f5gnssmf"
	// "github.com/p76081158/free5gc-nssmf"
	"github.com/p76081158/5g-nsmf/nsrtester/modules/nsrcreater"
	"github.com/p76081158/5g-nsmf/nsrtester/modules/yamltopara"
)

type DatasetInfo = yamltopara.DatasetInfo
type Resource = yamltopara.Resource

var Dir = "test-default"
var slcieRequestCase = "test-10"
var algorithm_list = []string {"right-top", "top-right", "leaf-size"}
var gnb_tenant_dictionary = []string {"466-01-000000010", "466-11-000000010", "466-93-000000010"}

var Tenant = 3
var SliceNum = 10
var ExtraRequestNum = 1
var TestNum = 5

var CpuLimit = 1000
var TimeWindowSize = 600
var ForecastBlockSize = 150
var Cpu_base = 100
var Cpu_min = 2
var Cpu_max = CpuLimit / Cpu_base
var Cpu_half = CpuLimit / (Cpu_base * 2)
var Cpu_lambda = Cpu_half
var Cpu_discount = 0.5

var BandwidthLimit = 10
var Bandwidth_min = 1
var Bandwidth_lambda = 5
var Bandwidth_discount = 0.5

var Slice_min_accept_num = 2
var Slice_duration = TimeWindowSize / Slice_min_accept_num

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <dir name> \n", os.Args[0])
		os.Exit(0)
	}
	if (os.Args[1]!="") {
		Dir = string(os.Args[1])
	}
	// get parameter from yaml file
	info := yamltopara.GetDataSetInfo(Dir)
	gnb_tenant_dictionary = info.NgciList
	Tenant = len(gnb_tenant_dictionary)
	SliceNum = info.SliceNum
	ExtraRequestNum = info.ExtraRequest
	TestNum = info.TestNum
	// cpu parameter
	CpuLimit = info.Resource.Cpu.Limit / Cpu_base
	ForecastBlockSize = info.ForecastBlockSize
	Cpu_lambda = info.Resource.Cpu.Lambda / Cpu_base
	Cpu_discount = info.Resource.Cpu.Discount
	// bandwidth parameter
	BandwidthLimit = info.Resource.Bandwidth.Limit
	Bandwidth_lambda = info.Resource.Bandwidth.Lambda
	Bandwidth_discount = info.Resource.Bandwidth.Discount
	// slice and timewindow info
	Slice_duration = info.Resource.Duration
	TimeWindowSize = info.Timewindow
	ForecastBlockSize = info.ForecastBlockSize
	fmt.Println(gnb_tenant_dictionary)
	// create dataset by parameter in yaml file
	nsrcreater.CreateDataSet(Dir, TestNum, gnb_tenant_dictionary, SliceNum, CpuLimit, Cpu_lambda, BandwidthLimit, Bandwidth_lambda, Slice_duration, ExtraRequestNum, ForecastBlockSize, Cpu_discount, Bandwidth_discount)
}