package nsrtoyaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/p76081158/5g-nsmf/modules/nsrhandler"
	"gopkg.in/yaml.v2"
)

type Yaml2GoRequestList      = nsrhandler.Yaml2GoRequestList
type Yaml2GoForecastingBlock = nsrhandler.Yaml2GoForecastingBlock
type RequestList             = nsrhandler.RequestList
type SliceList               = nsrhandler.SliceList
type ForecastingBlock        = nsrhandler.ForecastingBlock

// create dir for new set of network slice requests
func Mkdir(dir string) {
	sh_cmd    := "mkdir -p ../slice-requests/" + dir
	input_cmd := sh_cmd
	err       := exec.Command("/bin/sh", "-c", input_cmd).Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create network slice request set : ", dir)
}

func MkdirForecast(dir string) {
	sh_cmd    := "mkdir -p ../slice-forecasting/" + dir
	input_cmd := sh_cmd
	err       := exec.Command("/bin/sh", "-c", input_cmd).Run()
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

// wirte network slice request in the same timewindow to yaml
func WriteToXmlForecasting(src string, forecast []ForecastingBlock) {
	var snssai Yaml2GoForecastingBlock
	snssai.ForecastingBlock = forecast

	data,err := yaml.Marshal(&snssai)
	CheckError(err)
	err = ioutil.WriteFile(src,data,0777)
	CheckError(err)
}