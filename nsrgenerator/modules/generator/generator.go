package generator

import (
	"fmt"
	"math"
	"time"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/nsrtoyaml"
	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
)

type Slice            = slicebinpack.Slice
type Block            = slicebinpack.Block
type SliceList        = nsrhandler.SliceList
type ForecastingBlock = nsrhandler.ForecastingBlock

var Cpu_base = 100
var Cpu_min = 2
var Bandwidth_min = 1
var Slice_Duration_base = 100
var Slice_Duration_min = 1

// generate network slice info of each tenant and slice request
func SliceRequestGenerator(dir string, gnb_tenant_dictionary []string, num_slice int, cpu_max int, cpu_lambda int, bandwidthLimit int, bandwidth_lambda int, slice_duration int, slice_duration_random bool, timewindow_duration int, extra_request_num_each_timewindow int) {
	var slice_info_dictionary []SliceList
	num_tenant := len(gnb_tenant_dictionary)
	for i := 0; i < num_tenant; i++ {
		hex_tenant_index := fmt.Sprintf("%02x", i + 1)
		for j := 0; j < num_slice; j++ {
			r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
			cpu_poisson := distuv.Poisson{float64(cpu_lambda), r}
			bandwidth_poisson := distuv.Poisson{float64(bandwidth_lambda), r}
			hex_slice_index := fmt.Sprintf("%04x", j + 515)
			sliceCpu := int(cpu_poisson.Rand()) * Cpu_base
			slicebandwidth := int(bandwidth_poisson.Rand())
			sliceDuration := slice_duration
			if sliceCpu < (Cpu_min * Cpu_base) {
				sliceCpu = Cpu_min * Cpu_base
			} else if sliceCpu > (cpu_max * Cpu_base) {
				sliceCpu = cpu_max * Cpu_base
			}
			if slicebandwidth < Bandwidth_min {
				slicebandwidth = Bandwidth_min
			} else if slicebandwidth > bandwidthLimit {
				slicebandwidth = bandwidthLimit
			}
			if slice_duration_random {
				duration_lambda := slice_duration / Slice_Duration_base
				duration_poisson := distuv.Poisson{float64(duration_lambda), r}
				duration := int(duration_poisson.Rand())
				if duration < Bandwidth_min {
					duration = Bandwidth_min
				} else if duration > timewindow_duration / Slice_Duration_base {
					duration = timewindow_duration / Slice_Duration_base
				}
				sliceDuration = duration * Slice_Duration_base
			}

			s := SliceList {
				// snssai = sst(2bit) + sd(6bit)
				// sst = 01
				// sd first two bit defined by "hex_tenant_index"
				// sd last four bit defined by "hex_slice_index"
				// sd last four bit start by 0203
				// hex 0203 = dec 515
				Snssai:    "0x01" + hex_tenant_index + hex_slice_index,
				Ngci:      gnb_tenant_dictionary[i],
				Duration:  sliceDuration,
				Cpu:       sliceCpu,
				Bandwidth: slicebandwidth,
			}
			slice_info_dictionary = append(slice_info_dictionary, s)
		}
	}
	nsrtoyaml.WriteToXml("../slice-requests/" + dir + "/" + "slice-info-dictionary.yaml", slice_info_dictionary)
	RequsetGenerator(dir, slice_info_dictionary, num_tenant, num_slice, extra_request_num_each_timewindow)
}

// generate network slice request in each timewindow (dir of set, num of tenant, num of network slice each tenant, num of extra request)
// basic request number = tenant number
func RequsetGenerator(dir string, slice_info_dictionary []SliceList, num_tenant int, num_slice int, extra_request_num_each_timewindow int) {
	timewindow_num := int(math.Pow(float64(num_slice), float64(num_tenant)))
	for i := 0; i < timewindow_num; i++ {
		var requestSlices []SliceList
		// append by tenant id k
		for k := 0; k < num_tenant; k++ {
			tenant_base_index := k * num_slice
			if k == num_tenant - 1 {
				// slice_info_dictionary is one-dimension array, so need index transfer
				slice_index := (i % num_slice) + tenant_base_index
				requestSlices = append(requestSlices, slice_info_dictionary[slice_index])
			}else {                                                                                                           //        t1 t2 t3  (t = tenant)
				// slice_info_dictionary is one-dimension array, so need index transfer                                            e.g. 10 10 10  (num_slice)
				slice_index := (i / int(math.Pow(float64(num_slice), float64(num_tenant - k - 1)))) % 10 + tenant_base_index  //         0  0  0  =   0             t1_index =  i / 100
				requestSlices = append(requestSlices, slice_info_dictionary[slice_index])                                     //         0  0  1  =   1             t2_index =  i / 10
			}                                                                                                                 //         0  0  2  =   2             t3_index =  i % 10
		}                                                                                                                     //            .
		// append extra request                                                                                                          9  9  9  = 999 (i = 0~999)
		for j := 0; j < extra_request_num_each_timewindow; j++ {
			tenant_base_index := (i % num_tenant) * num_slice
			if (i % num_tenant) == num_tenant - 1 {
				// slice_info_dictionary is one-dimension array, so need index transfer
				slice_index := ((i % num_slice) + (num_slice - 1) * (i / num_tenant + 1)) % 10 + tenant_base_index
				requestSlices = append(requestSlices, slice_info_dictionary[slice_index])
			}else {
				// slice_info_dictionary is one-dimension array, so need index transfer
				slice_index := ((i / int(math.Pow(float64(num_slice),  float64(num_tenant - (i % num_tenant) - 1)))) + (num_slice - 1) * (i / num_tenant + 1)) % 10 + tenant_base_index
				requestSlices = append(requestSlices, slice_info_dictionary[slice_index])
			}
		}
		nsrtoyaml.WriteToXml("../slice-requests/" + dir + "/" + "timewindow-" + strconv.Itoa(i + 1) + ".yaml", requestSlices)	
	}
}

// generate forecaseting blocks for each slice 
func ForecaseGenerator(dir string, blcok_size int, cpu_max int, cpu_discount float64, bandwidthLimit int, bandwidth_discount float64) {
	slicelist := nsrtoyaml.GetSliceInfo(dir)
	for i := 0; i < len(slicelist); i++ {
		var blocklist []ForecastingBlock
		num_block := slicelist[i].Duration / blcok_size
		mkdir_folder := dir + "/" + slicelist[i].Ngci
		nsrtoyaml.MkdirForecast(mkdir_folder)
		src := "../slice-forecasting/" + mkdir_folder + "/" + slicelist[i].Snssai + ".yaml"

		for j := 0; j < num_block; j++ {
			r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
			cpu_resource := slicelist[i].Cpu / Cpu_base
			cpu_poisson := distuv.Poisson{float64(cpu_resource) * cpu_discount, r}
			blockCpu := int(cpu_poisson.Rand()) * Cpu_base
			bandwidth_resource := slicelist[i].Bandwidth
			bandwidth_poisson := distuv.Poisson{float64(bandwidth_resource) * bandwidth_discount, r}
			blockBandwidth := int(bandwidth_poisson.Rand())
			if blockCpu < (Cpu_min * Cpu_base) {
				blockCpu = Cpu_min * Cpu_base
			} else if blockCpu > (cpu_max * Cpu_base) {
				blockCpu = cpu_max * Cpu_base
			}
			if blockBandwidth < Bandwidth_min {
				blockBandwidth = Bandwidth_min
			} else if blockBandwidth > bandwidthLimit {
				blockBandwidth = bandwidthLimit
			}
			b := ForecastingBlock {
				Block:    j + 1,
				Duration: blcok_size,
				Cpu: blockCpu,
				Bandwidth: blockBandwidth,
			}
			blocklist = append(blocklist, b)
		}
		nsrtoyaml.WriteToXmlForecasting(src, blocklist)
	}
}