package nsrcreater

import (
	// "fmt"
	// "os"
	"strconv"

	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/nsrtoyaml"
	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/generator"
)

func CreateDataSet(dir string, test_num int, gnb_tenant_dictionary []string, sliceNum int, cpu_max int, cpu_lambda int, bandwidthLimit int, bandwidth_lambda int, slice_duration int, extra_request_num_each_timewindow int, forecastBlockSize int, cpu_discount float64, bandwidth_discount float64 ) {
	for i := 0; i < test_num; i++ {
		testName := "/test-" + strconv.Itoa(i + 1)
		path := dir + testName
		CreateTest(path, gnb_tenant_dictionary, sliceNum, cpu_max, cpu_lambda, bandwidthLimit, bandwidth_lambda, slice_duration, extra_request_num_each_timewindow, forecastBlockSize, cpu_discount, bandwidth_discount)
	}
}

func CreateTest(dir string, gnb_tenant_dictionary []string, sliceNum int, cpu_max int, cpu_lambda int, bandwidthLimit int, bandwidth_lambda int, slice_duration int, extra_request_num_each_timewindow int, forecastBlockSize int, cpu_discount float64, bandwidth_discount float64) {
	nsrtoyaml.Mkdir(dir)
	generator.SliceRequestGenerator(dir, gnb_tenant_dictionary, sliceNum, cpu_max, cpu_lambda, bandwidthLimit, bandwidth_lambda, slice_duration, extra_request_num_each_timewindow)
	generator.ForecaseGenerator(dir, forecastBlockSize, cpu_max, cpu_discount, bandwidthLimit, bandwidth_discount)
}