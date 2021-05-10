package main

import (
	"fmt"
	//"encoding/json"
	//"io/ioutil"

	"github.com/p76081158/5g-nsmf/slice_binpack"
)

// 1000 == 1000 milicore == i cpu
// 60 == 60 seconds
var ResourceLimit = 1000
var TimeWindowSize = 60

type Packer = slice_binpack.Packer
type Bin = slice_binpack.Bin
type Slice = slice_binpack.Slice
type Block = slice_binpack.Block
type SliceDeploy = slice_binpack.SliceDeploy

func main() {
	var request []Slice
	var bin Bin
	var access []Slice
	var reject []Slice
	var deploy_info []SliceDeploy
	var p Packer

	fmt.Println("")
	fmt.Println("Request1")
	fmt.Println("")
	// // input network slice requests
	request = append(request, Slice{"Slice 1", 30, 600, nil})
	request = append(request, Slice{"Slice 2", 20, 600, nil})
	request = append(request, Slice{"Slice 3", 20, 600, nil})
	
	bin = Bin{"Resource", 60, 1000, request}
	
	p = Packer{bin, access, reject, deploy_info}

	p.Pack()
	fmt.Println("Accept Slices: ", p.AcceptSlices)
	fmt.Println("Reject Slices: ", p.RejectSlices)
	fmt.Println("Deploy Info:   ", p.DeployInfos)

	fmt.Println("")
	fmt.Println("Request2")
	fmt.Println("")
	// input network slice requests
	request = request[:0]
	request = append(request, Slice{"Slice 1", 30, 600, nil})
	request = append(request, Slice{"Slice 2", 20, 600, []Block{{"1", 10, 600}, {"2", 10, 300}}})
	request = append(request, Slice{"Slice 3", 20, 600, nil})

	bin = Bin{"Resource", 60, 1000, request}

    p = Packer{bin, access, reject, deploy_info}
	p.Pack()
	fmt.Println("Accept Slices: ", p.AcceptSlices)
	fmt.Println("Reject Slices: ", p.RejectSlices)
	fmt.Println("Deploy Info:   ", p.DeployInfos)
}