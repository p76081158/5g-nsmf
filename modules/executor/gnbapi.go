package executor

import (
	"github.com/p76081158/5g-nsmf/modules/ueransim/gnb"
)

// get all gNB info. in Core Network
func GetgNBinfo() {
	gnb.GetgNBinfo()
}

// get all gNB and save as golang slice
func GetgNBlist() []string {
	return gnb.GetgNBlist()
}

// get all gNB and save as golang slice
func GetgNBdictionary() map[string]string {
	return gnb.GetgNBdictionary()
}

// get all gNB and save as golang slice
func GetgNB_B_dictionary() map[string]string {
	return gnb.GetgNB_B_dictionary()
}

// restart all gNB in Core Network
func RestartAllgNB() {
	gnb_list := GetgNBlist()
	for i := 0; i < len(gnb_list); i++ {
		gnb.RestartgNB(gnb_list[i])
	}
}