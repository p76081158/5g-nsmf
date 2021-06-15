package nsrtoyaml

import (
	"fmt"
	"io/ioutil"
	"log"
	// "os"
	"os/exec"
	// "strings"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"gopkg.in/yaml.v2"
)

type Yaml2GoRequestList = nsrhandler.Yaml2GoRequestList
type RequestList = nsrhandler.RequestList
type SliceList = nsrhandler.SliceList

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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// wirte network slice request in the same timewindow to yaml
func WriteToXml(src string, request []SliceList) {
	var timewindow Yaml2GoRequestList
	timewindow.RequestList.SliceList = request

	data,err := yaml.Marshal(&timewindow)
	CheckError(err)
	err = ioutil.WriteFile(src,data,0777)
	CheckError(err)
}