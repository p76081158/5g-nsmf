package nsrtoyaml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Get all network slice info
func GetSliceInfo(dir string) []SliceList {
    var timewindow    Yaml2GoRequestList
    var requestSlices []SliceList
    path          := "../slice-requests/" + dir + "/slice-info-dictionary.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &timewindow)
    if err != nil {
        panic(err)
    }

	requestSlices = timewindow.RequestList.SliceList
    return requestSlices
}