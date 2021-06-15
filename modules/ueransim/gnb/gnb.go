package gnb

import (
	"fmt"
	"log"
	"os/exec"
)

// restart gNB (ueransim-gnb will not work after long time idle)
func RestartgNB(ngci string) {
	gnb_cmd := "kubectl -n free5gc rollout restart deployment free5gc-ueransim-gnb-" + ngci
	input_cmd := gnb_cmd
	out, err := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(string(out))
}

// get all gNB custom resource info. in core network
func GetgNBinfo() {
	gnb_cmd := "kubectl -n free5gc get gnbs.nso.free5gc.com"
	input_cmd := gnb_cmd
	out, err := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Available gNB list in Core Network:")
	fmt.Println(string(out))
}

// create dictionary for gnb ip (static ip & n3 ip_B)
// func CreateIPDictionary() (map[string]string) {
// 	return
// }