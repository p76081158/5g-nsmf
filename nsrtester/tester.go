package main

import (
	// "fmt"
	// "time"
	// "strconv"

	// "github.com/p76081158/5g-nsmf/modules/nsrhandler"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	// "github.com/p76081158/5g-nsmf/modules/optimizer/scheduler"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/gnb"
	// "github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	// "github.com/p76081158/5g-nsmf/api/f5gnssmf"
	// "github.com/p76081158/free5gc-nssmf"
	"github.com/p76081158/5g-nsmf/nsrtest/modules/yamltopara"
)

type DatasetInfo = yamltopara.DatasetInfo
type Resource = yamltopara.Resource

var slcieRequestCase = "test-10"
var algorithm = "right-top"
// var timeWindowNumber = nsrhandler.GetTestCaseTimewindowNumber(slcieRequestCase)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <dir name> \n", os.Args[0])
		os.Exit(0)
	}
	if (os.Args[1]!="") {
		Dir = string(os.Args[1])
	}
	info := yamltopara.GetDataSetInfo(Dir)
	fmt.Println(info.Resource)
}