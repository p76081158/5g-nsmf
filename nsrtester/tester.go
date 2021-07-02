package main

import (
	"fmt"
	"os"

	"github.com/p76081158/5g-nsmf/nsrtester/modules/algorithmtester"
	"github.com/p76081158/5g-nsmf/nsrtester/modules/nsrcreater"
	"github.com/p76081158/5g-nsmf/nsrtester/modules/yamltopara"
)

type DatasetInfo = yamltopara.DatasetInfo
type Resource    = yamltopara.Resource

var Dir = "test-default"
var slcieRequestCase = "test-10"
var algorithm_list = []string {"pre-order", "invert-pre-order", "leaf-size"}
var gnb_tenant_dictionary = []string {"466-01-000000010", "466-11-000000010", "466-93-000000010"}

var Tenant = 3
var SliceNum = 10
var ExtraRequestNum = 1
var TestNum = 5
var RequestNum = Tenant + ExtraRequestNum

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
var Slice_duration_random = false
var TargetResource = "cpu"
// when use forecasting data, 0 == never, 1 == from first timewindow
var ForecastingTime = 1
var Sort = true
var Concat = false
var Regenerate = true

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
	RequestNum = Tenant + ExtraRequestNum
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
	Slice_duration_random = info.Resource.Random
	TargetResource = info.Target
	Sort = info.Sort
	Concat = info.Concat
	TimeWindowSize = info.Timewindow
	ForecastingTime = info.ForecastingTime
	ForecastBlockSize = info.ForecastBlockSize
	Regenerate = info.Regenerate
	fmt.Println(gnb_tenant_dictionary)
	// create dataset by parameter in yaml file	
	if Regenerate {
		nsrcreater.CreateDataSet(Dir, TestNum, gnb_tenant_dictionary, SliceNum, CpuLimit, Cpu_lambda, BandwidthLimit, Bandwidth_lambda, Slice_duration, Slice_duration_random, TimeWindowSize, ExtraRequestNum, ForecastBlockSize, Cpu_discount, Bandwidth_discount)
	}
	// test algorithm
	algorithmtester.TestAllAlgorithm(Dir, TestNum, algorithm_list, TargetResource, CpuLimit * Cpu_base, TimeWindowSize, RequestNum, ForecastingTime, Sort, Concat)
}