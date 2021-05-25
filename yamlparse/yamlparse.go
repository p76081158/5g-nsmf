package yamlparse

import (
    //"fmt"
    "io/ioutil"
    "log"
    "strings"
    "strconv"

    "github.com/p76081158/5g-nsmf/slice_binpack"
    "gopkg.in/yaml.v2"
)

type Slice = slice_binpack.Slice
type Block = slice_binpack.Block

func RefreshRequestList(windowID int, forecastingFinish bool) ([]Slice) {
    var timewindow Yaml2GoRequestList
    var requestSlices []Slice
    path := "slice-requests/timewindow-" + strconv.Itoa(windowID) + ".yaml"
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
        if forecastingFinish {
            sliceID := timewindow.RequestList.SliceList[i].Ngci + "," + timewindow.RequestList.SliceList[i].Snssai
            subBlock := GetForecastingBlock(sliceID)
            s := Slice {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Width:    timewindow.RequestList.SliceList[i].Duration,
                Height:   timewindow.RequestList.SliceList[i].Resource,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                SubBlock: subBlock,
            }
            requestSlices = append(requestSlices, s)
        } else {
            s := Slice {
                Name:     timewindow.RequestList.SliceList[i].Snssai,
                Width:    timewindow.RequestList.SliceList[i].Duration,
                Height:   timewindow.RequestList.SliceList[i].Resource,
                Ngci:     timewindow.RequestList.SliceList[i].Ngci,
                SubBlock: nil,
            }
            requestSlices = append(requestSlices, s)
        }
    }
    //fmt.Println(len(timewindow.RequestList.SliceList))
    return requestSlices
}

func GetForecastingBlock(sliceID string) ([]Block){
    var forecasting Yaml2GoForecastingBlock
    var requestBlock []Block
    split := strings.Split(sliceID, ",")
    ngci := split[0]
    slice := split[1]
    path := "slice-forecasting/" + ngci + "/" + slice + ".yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
        return nil
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
        requestBlock = append(requestBlock, b)
    }
    return requestBlock
}