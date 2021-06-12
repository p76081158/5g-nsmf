package generator

import (
	"fmt"
	"os/exec"
	"strconv"
)

type ResourcePattern struct {
	Resource int
	Duration int
}

type UeGenerator struct {
	Name     string
	Ngci     string
	RPs      []ResourcePattern
}

// create ueransim request-generator yaml file
func CreateGeneratorYaml(snssai string, ngci string, resourcePattern string, requestPattern string) {
	arg := snssai + " " + ngci + " " + resourcePattern + " " + requestPattern
	ue_cmd := "shell-script/ue-generator-create.sh " + arg
	input_cmd := ue_cmd
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	fmt.Println(resourcePattern)
	return
}

// 
func CreateUeRequestGenerator(ug []UeGenerator, requestPattern string) {
	for i := 0; i < len(ug); i++ {
		var resourcePattern string
		for j := 0; j < len(ug[i].RPs); j++ {
			if j != 0 {
				resourcePattern = resourcePattern + ","
			}
			resourcePattern = resourcePattern + strconv.Itoa(ug[i].RPs[j].Resource) + ":" + strconv.Itoa(ug[i].RPs[j].Duration)
		}
		fmt.Println(resourcePattern)
		fmt.Println(ug[i])
		CreateGeneratorYaml(ug[i].Name, ug[i].Ngci, resourcePattern, requestPattern)
	}
	return
}