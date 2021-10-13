package tenantbinpack

import (
	"fmt"
	"strconv"

	"github.com/p76081158/5g-nsmf/modules/executor"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
)

type Packer      = slicebinpack.Packer
type Bin         = slicebinpack.Bin
type SliceDeploy = slicebinpack.SliceDeploy
type Slice       = slicebinpack.Slice
type Block       = slicebinpack.Block
type DrawBlock   = slicebinpack.DrawBlock

var requestCpu       []Slice
var bin              Bin
var access           []Slice
var reject           []Slice
var deploy_info      []SliceDeploy
var draw_info        []DrawBlock
var p                Packer

var DrawScaleRatio      = 10

func RunTenant() {
	tenant1_access, tenant1_reject, tenant1_deploy := executor.RunTenantSliceBinPack("multi-tenant/466-01-000000010", "invert-pre-order")
	fmt.Println(tenant1_access)
	fmt.Println(tenant1_reject)
	fmt.Println(tenant1_deploy)
	tenant2_access, tenant2_reject, tenant2_deploy := executor.RunTenantSliceBinPack("multi-tenant/466-11-000000010", "invert-pre-order")
	fmt.Println(tenant2_access)
	fmt.Println(tenant2_reject)
	fmt.Println(tenant2_deploy)
	// fmt.Println("dsfdf")
	// fmt.Println(ToTenantItems_1("466-11-000000010", tenant2_reject, tenant2_deploy))
	tenant3_access, tenant3_reject, tenant3_deploy := executor.RunTenantSliceBinPack("multi-tenant/466-93-000000010", "invert-pre-order")
	fmt.Println(tenant3_access)
	fmt.Println(tenant3_reject)
	fmt.Println(tenant3_deploy)
	// fmt.Println("dsfdf")
	// fmt.Println(ToTenantItems_1("466-93-000000010", tenant3_reject, tenant3_deploy))

	requestCpu = append(requestCpu, ToTenantItems_1("466-01-000000010", tenant1_deploy)...)
	requestCpu = append(requestCpu, ToTenantItems_1("466-11-000000010", tenant2_deploy)...)
	requestCpu = append(requestCpu, ToTenantItems_1("466-93-000000010", tenant3_deploy)...)
	
	reject = append(reject, tenant1_reject...)
	reject = append(reject, tenant2_reject...)
	reject = append(reject, tenant3_reject...)
	if len(reject) > 0 {
		fmt.Println("tdsfdsfsdf")
		_, _, second_level_deploy := executor.RunTenantSliceBinPackByItem("multi-tenant/second-level-first-bin-pack", requestCpu, "invert-pre-order")
		item := ToTenantItems_1("second-level-compression", second_level_deploy)
		requestCpu = item
		requestCpu = append(requestCpu, reject...)
		bin        = Bin{"Resource", 600, 3000, requestCpu}
		p          = Packer{bin, access, reject, deploy_info, draw_info}
	
		p.Pack("invert-pre-order", true)
		// p.Pack("node-concat", true)
		
		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")
	
		// accept_count += len(p.AcceptSlices)
		// reject_count += len(p.RejectSlices)
		// scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
		slicebinpack.DrawBinPackResult("logs/binpack/" + "multi-tenant/", "1", p.DrawInfos, 600, 3000, DrawScaleRatio)
	} else {
		bin        = Bin{"Resource", 600, 3000, requestCpu}
		p          = Packer{bin, access, reject, deploy_info, draw_info}
	
		p.Pack("invert-pre-order", true)
		// p.Pack("node-concat", true)
		
		fmt.Println("Accept Slices: ", p.AcceptSlices)
		fmt.Println("Reject Slices: ", p.RejectSlices)
		fmt.Println("Deploy Info:   ", p.DeployInfos)
		fmt.Println("Draw Info:     ", p.DrawInfos)
		fmt.Println("")
	
		// accept_count += len(p.AcceptSlices)
		// reject_count += len(p.RejectSlices)
		// scheduler.SlicesScheduler(p.DeployInfos, gnb_ip_dictionary, gnb_ip_B_dictionary, DeployTimeBias, CPUofUserPlane, ueGenerator, RequestPattern)
		slicebinpack.DrawBinPackResult("logs/binpack/" + "multi-tenant/", "1", p.DrawInfos, 600, 3000, DrawScaleRatio)
	}




}

func ToTenantItems_1(tenant_name string, deploy []SliceDeploy) []Slice {
	// 0 100 200 300 400 500
	block        := []int{0, 0, 0, 0, 0, 0}
	blockList    := []Block{}
	newSliceList := []Slice{}
	max_width    := 0
	max_height   := 0
	for i := 0; i < len(deploy); i++ {
		for j := 0; j < deploy[i].Duration; j+=100 {
			block_index := ( deploy[i].Start + j ) / 100
			block[block_index] += deploy[i].Resource
			// fmt.Println(block_index)
		}
		// find max width and height
		if max_width < deploy[i].Start + deploy[i].Duration {
			max_width = deploy[i].Start + deploy[i].Duration
		}
		if max_height < deploy[i].Resource {
			max_height = deploy[i].Resource
		}
	}

	for i := 0; i < len(block); i++ {
		if block[i] > 0 {
			newBlock := Block {
				Name:     tenant_name + "-" + strconv.Itoa(i + 1),
				Width:    100,
				Height:   block[i],
			}
			blockList = append(blockList, newBlock)
		}
	}

	slice := Slice {
		Name:     tenant_name,
		Width:    max_width,
		Height:   max_height,
		Ngci:     tenant_name,
		SubBlock: blockList,
	}

	newSliceList = append(newSliceList, slice)
	return newSliceList
}

func ToTenantItems_2(tenant_name string, reject []Slice, deploy []SliceDeploy) []Slice {
	// 0 100 200 300 400 500
	block        := []int{0, 0, 0, 0, 0, 0}
	// blockList    := []Block{}
	blockList    := []Slice{}
	newSliceList := []Slice{}
	max_width    := 0
	max_height   := 0
	for i := 0; i < len(deploy); i++ {
		for j := 0; j < deploy[i].Duration; j+=100 {
			block_index := ( deploy[i].Start + j ) / 100
			block[block_index] += deploy[i].Resource
			// fmt.Println(block_index)
		}
		// find max width and height
		if max_width < deploy[i].Start + deploy[i].Duration {
			max_width = deploy[i].Start + deploy[i].Duration
		}
		if max_height < deploy[i].Resource {
			max_height = deploy[i].Resource
		}
	}

	for i := 0; i < len(block); i++ {
		if block[i] > 0 {
			// newBlock := Block {
			// 	Name:     tenant_name + "-" + strconv.Itoa(i + 1),
			// 	Width:    100,
			// 	Height:   block[i],
			// }
			newBlock := Slice {
				Name:     tenant_name + "-" + strconv.Itoa(i + 1),
				Width:    100,
				Height:   block[i],
				Ngci:     tenant_name,
				SubBlock: []Block{},
			}
			blockList = append(blockList, newBlock)
		}
	}

	// slice := Slice {
	// 	Name:     tenant_name,
	// 	Width:    max_width,
	// 	Height:   max_height,
	// 	Ngci:     tenant_name,
	// 	SubBlock: blockList,
	// }

	newSliceList = append(newSliceList, blockList...)
	newSliceList = append(newSliceList, reject...)
	return newSliceList
}