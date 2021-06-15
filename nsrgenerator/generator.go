package main

import (
	"fmt"
	// "log"
	"os"
	// "os/exec"
	"time"
	// "strings"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/nsrgenerator/modules/nsrtoyaml"
	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
)

type Slice = slicebinpack.Slice
type Block = slicebinpack.Block
type SliceList = nsrhandler.SliceList

var Dir = "test-default"
var Tenant = 3
var SliceNum = 10
var ResourceLimit = 1000
var TimeWindowSize = 600
var Cpu_min = 2
var Cpu_max = ResourceLimit / 100
var Cpu_half = ResourceLimit / 200
var Slice_min_accept_num = 2
var Slice_time = TimeWindowSize / Slice_min_accept_num
var gnb_tenant_dictionary = []string {"466-01-000000010", "466-11-000000010", "466-93-000000010"}

// generate network slice info of each tenant
func SliceInfoGenerator(dir string, num_tenant int, num_slice int) {
	var slice_info_dictionary []SliceList
	for i := 0; i < num_tenant; i++ {
		hex_tenant_index := fmt.Sprintf("%02x", i + 1)
		for j := 0; j < num_slice; j++ {
			r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
			resource_poisson := distuv.Poisson{float64(Cpu_half), r}
			// rand.Seed(uint64(time.Now().UnixNano()))
			// t := rand.Intn(Cpu_max - Cpu_min) + Cpu_min
			// fmt.Println(t)
			hex_slice_index := fmt.Sprintf("%04x", j + 515)
			sliceResource := int(resource_poisson.Rand()) * 100
			s := SliceList {
				// snssai = sst(2bit) + sd(6bit)
				// sst = 01
				// sd first two bit defined by "hex_tenant_index"
				// sd last four bit defined by "hex_slice_index"
				// sd last four bit start by 0203
				// hex 0203 = dec 515
				Snssai:   "0x01" + hex_tenant_index + hex_slice_index,
				Ngci:     gnb_tenant_dictionary[i],
				Duration: Slice_time,
				Resource: sliceResource,
			}
			slice_info_dictionary = append(slice_info_dictionary, s)
		}
	}
	nsrtoyaml.WriteToXml("../slice-requests/" + dir + "/" + "slice-info-dictionary.yaml", slice_info_dictionary)
}

// generate network slice request set (dir of set, timewindow num, )
func RequsetGenerator(dir string, timewindow_num int, request_num_each_timewindow int) {
	
	for i := 0; i < timewindow_num; i++ {
		var requestSlices []SliceList
		for j := 0; j < request_num_each_timewindow; j++ {
			s := SliceList {
				Snssai:   "0x01010203",
				Ngci:     "466-01-000000010",
				Duration: 300,
				Resource: 600,
			}
			requestSlices = append(requestSlices, s)
		}
		nsrtoyaml.WriteToXml("../slice-requests/" + dir + "/" + "timewindow-" + strconv.Itoa(i + 1) + ".yaml", requestSlices)
	}
}

func handleError(err error) {
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
}

func main() {
	if len(os.Args) != 6 {
		fmt.Printf("Usage : %s <dir name> <tenant number> <slice number of each tenant> <limit of cpu resource> <size of each timewindow>\n", os.Args[0])
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
		ResourceLimit = i
	}
	if (os.Args[5]!="") {
		i, err := strconv.Atoi(string(os.Args[5]))
		handleError(err)
		TimeWindowSize = i
	}

	nsrtoyaml.Mkdir(Dir)
	SliceInfoGenerator(Dir, Tenant, SliceNum)
	// RequsetGenerator(Dir)
	
	// fmt.Println(hex)
}