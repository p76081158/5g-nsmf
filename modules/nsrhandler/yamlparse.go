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
func GetSliceInfo(dir string) []Slice {
    var timewindow Yaml2GoRequestList
    var requestSlices []Slice
    path := "slice-requests/" + dir + "/slice-info-dictionary.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &timewindow)
    if err != nil {
        panic(err)
    }

    slice_num := len(timewindow.RequestList.SliceList)
    for i := 0; i < slice_num; i++ {
        s := Slice {
            Name:     timewindow.RequestList.SliceList[i].Snssai,
            Width:    timewindow.RequestList.SliceList[i].Duration,
            Height:   timewindow.RequestList.SliceList[i].Resource,
            Ngci:     timewindow.RequestList.SliceList[i].Ngci,
            SubBlock: nil,
        }
        requestSlices      = append(requestSlices, s)

    }
    return requestSlices
}

// get network slice requests base on test case dir and time_window_id, alse generate ue request pattern
func RefreshRequestList(dir string, windowID int, forecastingFinish bool) ([]Slice, []UeGenerator) {
    var timewindow Yaml2GoRequestList
    var requestSlices []Slice
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
    sort.Sort(ByResource(timewindow.RequestList.SliceList))
    slice_num := len(timewindow.RequestList.SliceList)
    for i := 0; i < slice_num; i++ {
        if forecastingFinish {
            sliceID := timewindow.RequestList.SliceList[i].Ngci + "," + timewindow.RequestList.SliceList[i].Snssai
            subBlock, rps := GetForecastingBlock(dir, sliceID)
            if rps == nil {
                rps = []ResourcePattern{}
                rps = append(rps, ResourcePattern {
                    Resource:  timewindow.RequestList.SliceList[i].Resource,
                    Duration:  timewindow.RequestList.SliceList[i].Duration,
                })
            }
            s := Slice {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Width:    timewindow.RequestList.SliceList[i].Duration,
                Height:   timewindow.RequestList.SliceList[i].Resource,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                SubBlock: subBlock,
            }
            u := UeGenerator {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                RPs:      rps,
            }
            requestSlices      = append(requestSlices, s)
            requestUeGenerator = append(requestUeGenerator, u)
        } else {
            rps := []ResourcePattern{}
            s := Slice {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Width:    timewindow.RequestList.SliceList[i].Duration,
                Height:   timewindow.RequestList.SliceList[i].Resource,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                SubBlock: nil,
            }
            u := UeGenerator {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                RPs:      append(rps, ResourcePattern {
                    Resource:  timewindow.RequestList.SliceList[i].Resource,
                    Duration:  timewindow.RequestList.SliceList[i].Duration,
                }),
            }
            requestSlices      = append(requestSlices, s)
            requestUeGenerator = append(requestUeGenerator, u)
        }
    }
    return requestSlices, requestUeGenerator
}

// get forecasted network slice and ue request pattern
func GetForecastingBlock(dir string, sliceID string) ([]Block, []ResourcePattern) {
    var forecasting Yaml2GoForecastingBlock
    var requestBlock []Block
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
        b := Block {
            Name:     slice + "-" + strconv.Itoa(forecasting.ForecastingBlock[i].Block),
            Width:    forecasting.ForecastingBlock[i].Duration,
            Height:   forecasting.ForecastingBlock[i].Resource,
        }
        r := ResourcePattern {
            Resource: forecasting.ForecastingBlock[i].Resource,
            Duration: forecasting.ForecastingBlock[i].Duration,
        }
        requestBlock   = append(requestBlock, b)
        resourcePattern = append(resourcePattern, r)
    }
    return requestBlock, resourcePattern
}