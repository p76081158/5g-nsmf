package yamlparse

import (
    "fmt"
    "io/ioutil"
    "log"
    //"path/filepath"
    //"os"
    "strconv"

    "github.com/p76081158/5g-nsmf/slice_binpack"
    "gopkg.in/yaml.v2"
)

// func (c *Timewindow) getConf() *Timewindow {

//     yamlFile, err := ioutil.ReadFile("slice-requests/timewindow-1.yaml")
//     if err != nil {
//         log.Printf("yamlFile.Get err   #%v ", err)
//     }
//     err = yaml.Unmarshal(yamlFile, c)
//     if err != nil {
//         log.Fatalf("Unmarshal: %v", err)
//     }

//     return c
// }


// func Test() {
//     var qq Timewindow
//     qq.getConf()
//     fmt.Println(qq)
// }
type Slice = slice_binpack.Slice

func RefreshRequestList(windowID int, forecastingFinish bool) ([]Slice) {
    var timewindow Yaml2Go
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
            // not done
            // s := Slice {
            //     Name:     timewindow.RequestList.SliceList[i].Snssai,
            //     Width:    timewindow.RequestList.SliceList[i].Duration,
            //     Height:   timewindow.RequestList.SliceList[i].Resource,
            //     Ngci:     timewindow.RequestList.SliceList[i].Ngci,
            //     SubBlock: ,
            // }
            // requestSlices = append(request, s)
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
    fmt.Println(len(timewindow.RequestList.SliceList))
    return requestSlices
}




// func GetRequestList(timewindowid int) *Timewindow {
//     path := "slice-requests/timewindow-" + strconv.Itoa(timewindowid)  + ".yaml"
//     filename, _ := filepath.Abs(path)
//     yamlFile, err := ioutil.ReadFile(filename)

//     if err != nil {
//         panic(err)
//     }

//     var timewindow Timewindow

//     err = yaml.Unmarshal(yamlFile, &timewindow)
//     if err != nil {
//         panic(err)
//     }

//     return timewindow
// }
// var TimewindowNow Timewindow
// var TimewindowNowRequest Request

// func RefreshRequestList(timewindowid int)  {
// 	path := "slice-requests/timewindow-" + strconv.Itoa(timewindowid)  + ".yaml"
//     fmt.Println(path)
//     yamlFile, err := ioutil.ReadFile(path)

//     if err != nil {
//         panic(err)
//     }

//     TimewindowNow = Timewindow{}

//     err = yaml.Unmarshal(yamlFile, &TimewindowNow)
//     if err != nil {
//         panic(err)
//     }
//     var test Request
//     test = *(TimewindowNow.Request)
//     fmt.Println(test)
// }

// func readConf(timewindowid int)) (*myData, error) {

//     buf, err := ioutil.ReadFile(filename)
//     if err != nil {
//         return nil, err
//     }

//     c := &myData{}
//     err = yaml.Unmarshal(buf, c)
//     if err != nil {
//         return nil, fmt.Errorf("in file %q: %v", filename, err)
//     }

//     return c, nil
// }