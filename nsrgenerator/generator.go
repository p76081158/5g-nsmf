package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"io/ioutil"
    // "strings"
    "strconv"

	"github.com/p76081158/5g-nsmf/module/nsrhandler"
	"github.com/p76081158/5g-nsmf/module/optimizer/slicebinpack"
    "gopkg.in/yaml.v2"
	//"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
)

type Slice = slicebinpack.Slice
type Block = slicebinpack.Block
type Yaml2GoRequestList = nsrhandler.Yaml2GoRequestList
type RequestList = nsrhandler.RequestList
type SliceList = nsrhandler.SliceList

var Dir = "test-1"
var Tenant = 3
var Num = 10
var Cpu_min = 2
var Cpu_max = 10


// create dir for new set of network slice requests
func Mkdir(dir string) {
	sh_cmd := "mkdir -p ../slice-requests/" + dir
	input_cmd := sh_cmd
	err := exec.Command("/bin/sh", "-c", input_cmd).Run()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("create network slice request set : ", dir)
}



// generate network slice request set (dir of set, timewindow num, )
func RequsetGenerator(dir string, ) {
	
	for i := 0; i < n; i++ {
		var requestSlices []SliceList
		for j := 0; j < k; j++ {
			s := SliceList {
				Snssai:   "0x01010203",
				Ngci:     "466-01-000000010",
				Duration: 300,
				Resource: 600,
			}
			requestSlices = append(requestSlices, s)
		}
		writeToXml("../slice-requests/" + dir + "/" + "timewindow-" + i + ".yaml", requestSlices)
	}
	
}

func checkError(err error) {
	if err != nil {
        panic(err)
    }
}

// wirte network slice request in the same timewindow to yaml
func writeToXml(src string, request []SliceList) {
	var timewindow Yaml2GoRequestList
    timewindow.RequestList.SliceList = request

	data,err := yaml.Marshal(&timewindow)
	checkError(err)
	err = ioutil.WriteFile(src,data,0777)
	checkError(err)
}

func main() {
	// if len(os.Args) != 5 {
	// 	fmt.Printf("Usage : %s <curl cmd> <specify interface name> <resource pattern> <request ratio>\n", os.Args[0])
	// 	os.Exit(0)
	// }
	if (os.Args[1]!="") {
		Dir = string(os.Args[1])
	}
	if (os.Args[2]!="") {
		i, err := strconv.Atoi(string(os.Args[2]))
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		Num = i
	}
	if (os.Args[3]!="") {
		i, err := strconv.Atoi(string(os.Args[3]))
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		Tenant = i
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	t := rand.Intn(Cpu_max - Cpu_min) + Cpu_min
	Mkdir(Dir)
	RequsetGenerator(Dir)
	fmt.Println(t)
}