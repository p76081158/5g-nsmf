package algorithmtester

import (
	"fmt"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/nsrtester/modules/excelwriter"
)

// struct alias
type Packer = slicebinpack.Packer
type Bin = slicebinpack.Bin
type Slice = slicebinpack.Slice
type Block = slicebinpack.Block
type SliceDeploy = slicebinpack.SliceDeploy
type DrawBlock = slicebinpack.DrawBlock

var caseCount []int
var access []Slice
var reject []Slice
var deploy_info []SliceDeploy
var draw_info []DrawBlock
var DrawScaleRatio = 10

// create golang slice for store accept case
func InitAcceptCase(caseNum int) {
	caseCount = []int{}
	for i := 0; i < caseNum; i++ {
		caseCount = append(caseCount, 0)
	}
}

// get string of slice which stores accept case list
func AcceptCaseToString() []string {
	var result []string
	for i := 0; i < len(caseCount); i++ {
		result = append(result, strconv.Itoa(caseCount[i]))
	}
	return result
}

// count number of each accept case
func CountAcceptCase(accept_list []Slice) {
	if len(accept_list) != 0 {
		caseIndex := len(accept_list) - 1
		caseCount[caseIndex]++
	}
}

// create csv title
func InitCsvRowTitle(algorithm_name string, caseNum int) [][]string {
	var result_csv [][]string
	row_title := []string{"Algorithm", algorithm_name}
	test_case_title := []string{"name", "accept rate"}
	for i := 0; i < caseNum; i++ {
		accept_case_name := "accept-" + strconv.Itoa(i + 1)
		test_case_title = append(test_case_title, accept_case_name)
	}
	result_csv = append(result_csv, row_title)
	result_csv = append(result_csv, test_case_title)
	return result_csv
}

// test algorithm on each test cases in the dataset
func TestAlgorithm(dataset string, test_num int, algorithm string, resource string, resource_limit int, timewindowSize int, request_in_timewindow int, forecastingTime int, sort bool) [][]string {
	result_csv := InitCsvRowTitle(algorithm, request_in_timewindow)
	for i := 0; i < test_num; i++ {
		var test_case_csv []string
		forecastingFinish := false
		test_name := "test-" + strconv.Itoa(i + 1)
		path := "../slice-requests/" + dataset + "/" + test_name
		path_forecasting := "../slice-forecasting/" + dataset + test_name
		drawpath := "../logs/binpack/" + dataset + "/" + test_name + "/" + resource + "/" + algorithm
		slicebinpack.Mkdir(drawpath)
		timeWindowNumber := nsrhandler.GetTestCaseTimewindowNumber(path)
		accept_count := 0
		reject_count := 0
		InitAcceptCase(request_in_timewindow)
		// read every timewindow inside the test case
		for j := 0; j < timeWindowNumber; j++ {
			var bin Bin
			count := j + 1
			// use forecasting data or not
			if count == forecastingTime {
				forecastingFinish = true
			}
			// get network slice request by test case, and generate ue request pattern by network slice request
			requestCpu, requestBandwidth, _ := nsrhandler.RefreshRequestList(path, path_forecasting, count, forecastingFinish, sort)
			
			// select cpu or bandwidth
			if resource == "cpu" {
				bin = Bin{"Resource", timewindowSize, resource_limit, requestCpu}
			} else if resource == "bandwidth" {
				bin = Bin{"Resource", timewindowSize, resource_limit, requestBandwidth}
			}
			p := Packer{bin, access, reject, deploy_info, draw_info}
			p.Pack(algorithm)
	
			fmt.Println("Accept Slices: ", p.AcceptSlices)
			fmt.Println("Reject Slices: ", p.RejectSlices)
			fmt.Println("Deploy Info:   ", p.DeployInfos)
			fmt.Println("Draw Info:     ", p.DrawInfos)
			fmt.Println("")
	
			accept_count += len(p.AcceptSlices)
			reject_count += len(p.RejectSlices)
			CountAcceptCase(p.AcceptSlices)
			// scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
			slicebinpack.DrawBinPackResult(drawpath, strconv.Itoa(count), p.DrawInfos, timewindowSize, resource_limit, DrawScaleRatio)
		}
		accept_rate := float64(accept_count) / float64( request_in_timewindow *  timeWindowNumber )
		string_accept_rate := fmt.Sprintf("%f", accept_rate)
		test_case_csv = append(test_case_csv, test_name)
		test_case_csv = append(test_case_csv, string_accept_rate)
		test_case_csv = append(test_case_csv, AcceptCaseToString()...)
		result_csv    = append(result_csv, test_case_csv)
		fmt.Println("Accept count: ", accept_count)
		fmt.Println("Reject count: ", reject_count)
		fmt.Println("Accept rate : ", accept_rate)
	}
	return result_csv
}

// test all algorithm in algorithm_list
func TestAllAlgorithm(dataset string, test_num int, algorithm_list []string, resource string, resource_limit int, timewindowSize int, request_in_timewindow int, forecastingTime int, sort bool) {
	var csv_data [][]string
	for i := 0; i < len(algorithm_list); i++ {
		result := TestAlgorithm(dataset, test_num, algorithm_list[i], resource, resource_limit, timewindowSize, request_in_timewindow, forecastingTime, sort)
		csv_data = append(csv_data, result...)
	}
	excelwriter.WriteToExcel(dataset, strconv.Itoa(forecastingTime), csv_data)
}