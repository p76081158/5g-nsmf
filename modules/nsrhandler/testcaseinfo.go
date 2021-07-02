package nsrhandler

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// get number of Timewindow in specific test case
func GetTestCaseTimewindowNumber(case_dir string) int {
	test_case_cmd := "ls " + case_dir + "/ | grep -c timewindow-"
	input_cmd     := test_case_cmd
	out, err      := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	temp          := strings.Split(string(out), "\n")
	length, err   := strconv.Atoi(temp[0])
	return length
}