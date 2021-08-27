package gnb

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// restart gNB (ueransim-gnb will not work after long time idle)
func RestartgNB(ngci string) {
	gnb_cmd   := "kubectl -n free5gc rollout restart deployment free5gc-ueransim-gnb-" + ngci
	input_cmd := gnb_cmd
	out, err  := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(string(out))
}

// get all gNB custom resource info. in core network
func GetgNBinfo() {
	gnb_cmd   := "kubectl -n free5gc get gnbs.nso.free5gc.com"
	input_cmd := gnb_cmd
	out, err  := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Available gNB list in Core Network:")
	fmt.Println(string(out))
}

// get gNB list and save as golang slice
func GetgNBlist() []string {
	gnb_cmd   := "kubectl -n free5gc get gnbs.nso.free5gc.com | awk -F ' ' '{print $1}'"
	input_cmd := gnb_cmd
	out, err  := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Available gNB list in Core Network:")
	fmt.Println(string(out))
	gnb_list := strings.Split(string(out), "\n")
	return gnb_list[1:len(gnb_list)-1]
}

// create dictionary for gnb ip (static ip & n3 ip_B)
// func CreateIPDictionary() (map[string]string) {
// 	return
// }