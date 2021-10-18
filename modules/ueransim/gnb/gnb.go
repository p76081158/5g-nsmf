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

// (base) vcx@vcx-virtual-machine:~/5g-nsmf$ kubectl -n free5gc get gnbs.nso.free5gc.com
// NAME               MCC   MNC   UE NUMBERS   N3 CIDR           EXTERNAL IP
// 466-01-000000010   466   01    0            10.200.100.0/24   192.168.72.51
// 466-11-000000010   466   11    0            10.201.100.0/24   192.168.72.53
// 466-93-000000010   466   93    0            10.202.100.0/24   192.168.72.55

// get gNB dictionary and save as golang slice
func GetgNBdictionary() map[string]string {
	gnb_cmd   := "kubectl -n free5gc get gnbs.nso.free5gc.com | awk -F ' ' '{print $1,$6}'"
	input_cmd := gnb_cmd
	out, err  := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Available gNB list in Core Network:")
	fmt.Println(string(out))
	gnb_list          := strings.Split(string(out), "\n")
	gnb_list           = gnb_list[1:len(gnb_list)-1]
	gnb_ip_dictionary := make(map[string]string)

	for i := 0; i < len(gnb_list); i++ {
		gnb_dict                      := strings.Split(gnb_list[i], " ")
		gnb_ip_dictionary[gnb_dict[0]] = gnb_dict[1]
	}

	// gnb_ip_dictionary := map[string]string{"466-01-000000010": "192.168.72.51", "466-11-000000010": "192.168.72.53", "466-93-000000010": "192.168.72.55"}
	return gnb_ip_dictionary
}

// CIDR == A.B.C.D
// get gNB B dictionary and save as golang slice (N3 ip mapping)
func GetgNB_B_dictionary() map[string]string {
	gnb_cmd   := "kubectl -n free5gc get gnbs.nso.free5gc.com | awk -F ' ' '{print $1,$5}'"
	input_cmd := gnb_cmd
	out, err  := exec.Command("/bin/sh", "-c", input_cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Available gNB list in Core Network:")
	fmt.Println(string(out))
	gnb_list            := strings.Split(string(out), "\n")
	gnb_list             = gnb_list[1:len(gnb_list)-1]
	gnb_ip_B_dictionary := make(map[string]string)

	for i := 0; i < len(gnb_list); i++ {
		gnb_dict                        := strings.Split(gnb_list[i], " ")
		gnb_b                           := strings.Split(gnb_dict[1], ".")
		gnb_ip_B_dictionary[gnb_dict[0]] = gnb_b[1]
	}

	// gnb_ip_B_dictionary = map[string]string{"466-01-000000010": "200", "466-11-000000010": "201", "466-93-000000010": "202"}
	return gnb_ip_B_dictionary
}

// create dictionary for gnb ip (static ip & n3 ip_B)
// func CreateIPDictionary() (map[string]string) {
// 	return
// }