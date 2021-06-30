package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/nsrtoyaml"
	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/generator"
)

var Dir = "test-default"
var Tenant = 3
var SliceNum = 10
var ExtraSliceNum = 1
var CpuLimit = 1000
var TimeWindowSize = 600
var ForecastBlockSize = 150
var Cpu_base = 100
var Cpu_max = CpuLimit / Cpu_base
var Cpu_half = CpuLimit / (Cpu_base * 2)
var Cpu_lambda = Cpu_half
var Cpu_discount = 0.5
var BandwidthLimit = 10
var Bandwidth_lambda = 5
var Bandwidth_discount = 0.5
var Slice_min_accept_num = 2
var Slice_duration = TimeWindowSize / Slice_min_accept_num
var Slice_duration_random = false
var gnb_tenant_dictionary = []string {"466-01-000000010", "466-11-000000010", "466-93-000000010"}

func handleError(err error) {
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
}

func main() {
	if len(os.Args) != 11 {
		fmt.Printf("Usage : %s <dir name> <tenant number> <slice number of each tenant> <limit of cpu resource> <lambda of cpu resource> <discount of cpu resource> <limit of bandwidth resource> <lambda of bandwidth resource> <discount of bandwidth resource> <size of each timewindow>\n", os.Args[0])
		os.Exit(0)
	}
	if (os.Args[1]!="") {
		Dir = string(os.Args[1])
	}
	if (os.Args[2]!="") {
		i, err := strconv.Atoi(string(os.Args[2]))
		handleError(err)
		Tenant = i
	}
	if (os.Args[3]!="") {
		i, err := strconv.Atoi(string(os.Args[3]))
		handleError(err)
		SliceNum = i
	}
	if (os.Args[4]!="") {
		i, err := strconv.Atoi(string(os.Args[4]))
		handleError(err)
		CpuLimit = i
	}
	if (os.Args[5]!="") {
		i, err := strconv.Atoi(string(os.Args[5]))
		handleError(err)
		Cpu_lambda = i
	}
	if (os.Args[6]!="") {
		i, err := strconv.ParseFloat(string(os.Args[6]), 64)
		handleError(err)
		Cpu_discount = i
	}
	if (os.Args[7]!="") {
		i, err := strconv.Atoi(string(os.Args[7]))
		handleError(err)
		BandwidthLimit = i
	}
	if (os.Args[8]!="") {
		i, err := strconv.Atoi(string(os.Args[8]))
		handleError(err)
		Bandwidth_lambda = i
	}
	if (os.Args[9]!="") {
		i, err := strconv.ParseFloat(string(os.Args[9]), 64)
		handleError(err)
		Bandwidth_discount = i
	}
	if (os.Args[10]!="") {
		i, err := strconv.Atoi(string(os.Args[10]))
		handleError(err)
		TimeWindowSize = i
	}

	nsrtoyaml.Mkdir(Dir)
	generator.SliceRequestGenerator(Dir, gnb_tenant_dictionary, SliceNum, Cpu_max, Cpu_lambda, BandwidthLimit, Bandwidth_lambda, Slice_duration, Slice_duration_random, TimeWindowSize, ExtraSliceNum)
	generator.ForecaseGenerator(Dir, ForecastBlockSize, Cpu_max, Cpu_discount, BandwidthLimit, Bandwidth_discount)
}

// test example:
// ./nsrgenerator DataSet-test/test1 3 10 1000 5 0.5 10 5 0.5 600