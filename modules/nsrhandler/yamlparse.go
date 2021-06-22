package nsrhandler

import (
    "io/ioutil"
    "log"
    "sort"
    "strings"
    "strconv"

    "github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
    "github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
    "gopkg.in/yaml.v2"
)

type Slice = slicebinpack.Slice
type Block = slicebinpack.Block
type ResourcePattern = generator.ResourcePattern
type UeGenerator = generator.UeGenerator

// Get all network slice info
func GetSliceInfo(dir string) []SliceList {
    var timewindow Yaml2GoRequestList
    var requestSlicesList []SliceList
    path := "slice-requests/" + dir + "/slice-info-dictionary.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &timewindow)
    if err != nil {
        panic(err)
    }
    requestSlicesList = timewindow.RequestList.SliceList
    return requestSlicesList
}

// get network slice requests base on test case dir and time_window_id, alse generate ue request pattern
func RefreshRequestList(dir string, windowID int, forecastingFinish bool) ([]Slice, []Slice, []UeGenerator) {
    var timewindow Yaml2GoRequestList
    var slicelist_cpu []SliceList
    var slicelist_bandwidth []SliceList
    var requestSlicesCpu []Slice
    var requestSlicesBandwidth []Slice
    var requestUeGenerator []UeGenerator
    path := "slice-requests/" + dir + "/timewindow-" + strconv.Itoa(windowID) + ".yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &timewindow)
    if err != nil {
        panic(err)
    }

    // sort by Resource
    slicelist_cpu       = timewindow.RequestList.SliceList
    slicelist_bandwidth = timewindow.RequestList.SliceList
    sort.Sort(ByCpu(slicelist_cpu))
    sort.Sort(ByBandwidth(slicelist_bandwidth))
    
    slice_num := len(timewindow.RequestList.SliceList)
    for i := 0; i < slice_num; i++ {
        if forecastingFinish {
            sliceID := timewindow.RequestList.SliceList[i].Ngci + "," + timewindow.RequestList.SliceList[i].Snssai
            subBlockCpu, subBlockBandwidth, rps := GetForecastingBlock(dir, sliceID)
            if rps == nil {
                rps = []ResourcePattern{}
                rps = append(rps, ResourcePattern {
                    Resource:  slicelist_cpu[i].Cpu,
                    Duration:  slicelist_cpu[i].Duration,
                })
            }
            s_cpu := Slice {
                Name:     slicelist_cpu[i].Snssai,
                Width:    slicelist_cpu[i].Duration,
                Height:   slicelist_cpu[i].Cpu,
                Ngci:     slicelist_cpu[i].Ngci,
                SubBlock: subBlockCpu,
            }
            s_bandwidth := Slice {
                Name:     slicelist_bandwidth.Snssai,
                Width:    slicelist_bandwidth.Duration,
                Height:   slicelist_bandwidth.Bandwidth,
                Ngci:     slicelist_bandwidth.Ngci,
                SubBlock: subBlockBandwidth,
            }
            u := UeGenerator {
                Name:     slicelist_cpu[i].Snssai,
                Ngci:     slicelist_cpu[i].Ngci,
                RPs:      rps,
            }
            requestSlicesCpu       = append(requestSlicesCpu, s_cpu)
            requestSlicesBandwidth = append(requestSlicesBandwidth, s_bandwidth)
            requestUeGenerator     = append(requestUeGenerator, u)
        } else {
            rps := []ResourcePattern{}
            s_cpu := Slice {
                Name:     slicelist_cpu[i].Snssai,
                Width:    slicelist_cpu[i].Duration,
                Height:   slicelist_cpu[i].Cpu,
                Ngci:     slicelist_cpu[i].Ngci,
                SubBlock: nil,
            }
            s_bandwidth := Slice {
                Name:     slicelist_bandwidth.Snssai,
                Width:    slicelist_bandwidth.Duration,
                Height:   slicelist_bandwidth.Bandwidth,
                Ngci:     slicelist_bandwidth.Ngci,
                SubBlock: nil,
            }
            u := UeGenerator {
                Name:     slicelist_cpu[i].Snssai,
                Ngci:     slicelist_cpu[i].Ngci,
                RPs:      append(rps, ResourcePattern {
                    Resource:  slicelist_cpu[i].Cpu,
                    Duration:  slicelist_cpu[i].Duration,
                }),
            }
            requestSlicesCpu       = append(requestSlicesCpu, s_cpu)
            requestSlicesBandwidth = append(requestSlicesBandwidth, s_bandwidth)
            requestUeGenerator     = append(requestUeGenerator, u)
        }
    }
    return requestSlicesCpu, requestSlicesBandwidth, requestUeGenerator
}

// get forecasted network slice and ue request pattern
func GetForecastingBlock(dir string, sliceID string) ([]Block, []Block, []ResourcePattern) {
    var forecasting Yaml2GoForecastingBlock
    var requestBlockCpu []Block
    var requestBlockBandwidth []Block
    var resourcePattern []ResourcePattern
    split := strings.Split(sliceID, ",")
    ngci := split[0]
    slice := split[1]
    path := "slice-forecasting/" + dir + "/" + ngci + "/" + slice + ".yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
        return nil, nil
    }
    err = yaml.Unmarshal(yamlFile, &forecasting)
    if err != nil {
        panic(err)
    }
    block_num := len(forecasting.ForecastingBlock)
    for i := 0; i < block_num; i++ {
        b_cpu := Block {
            Name:     slice + "-" + strconv.Itoa(forecasting.ForecastingBlock[i].Block),
            Width:    forecasting.ForecastingBlock[i].Duration,
            Height:   forecasting.ForecastingBlock[i].Cpu,
        }
        b_bandwidth := Block {
            Name:     slice + "-" + strconv.Itoa(forecasting.ForecastingBlock[i].Block),
            Width:    forecasting.ForecastingBlock[i].Duration,
            Height:   forecasting.ForecastingBlock[i].Bandwidth,
        }
        // ue cpu resource pattern
        r := ResourcePattern {
            Resource: forecasting.ForecastingBlock[i].Cpu,
            Duration: forecasting.ForecastingBlock[i].Duration,
        }
        requestBlockCpu       = append(requestBlockCpu, b_cpu)
        requestBlockBandwidth = append(requestBlockBandwidth, b_bandwidth)
        resourcePattern       = append(resourcePattern, r)
    }
    return requestBlockCpu, requestBlockBandwidth, resourcePattern
}