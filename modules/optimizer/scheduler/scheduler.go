package scheduler

import (
	"strings"

	// "fmt"
	"github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack"
	"github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator"
	"github.com/p76081158/5g-nsmf/api/f5gnssmf"
)

type SliceDeploy = slicebinpack.SliceDeploy
type UeGenerator = generator.UeGenerator

// schedule network slices in the same time_window
func SlicesScheduler(slicesDeploy []SliceDeploy, gnb_ip_dictionary map[string]string, gnb_ip_B_dictionary map[string]string, deploy_time_bias float64, cpu_of_user_plane int, ue_generator []UeGenerator, requestPattern string) {
	generator.CreateUeRequestGenerator(ue_generator, requestPattern)
	for i :=0; i< len(slicesDeploy); i++ {
		slice_name  := strings.Split(slicesDeploy[i].Name, "-")
		gnb_ip      := gnb_ip_dictionary[slicesDeploy[i].Ngci]
		gnb_n3_ip_B := gnb_ip_B_dictionary[slicesDeploy[i].Ngci]
		ngci        := slicesDeploy[i].Ngci
		start       := slicesDeploy[i].Start
		duration    := slicesDeploy[i].Duration
		end         := slicesDeploy[i].End
		cpu         := slicesDeploy[i].Resource

		
		// modify existed network slice or create new network slice
		if len(slice_name) > 1 && slice_name[1] != "1" {
			//go SliceModifyServiceCPU(slice_name[0], ngci, cpu, start, duration, end)
			go f5gnssmf.SliceModifyServiceCPU(slice_name[0], ngci, cpu, start, duration, end)
		} else {
			//go ApplySliceToCoreNetwork(slice_name[0], gnb_ip, gnb_n3_ip_B, ngci, cpu, CPUofUserPlane, start, duration, end)
			go f5gnssmf.ApplySliceToCoreNetwork(slice_name[0], gnb_ip, gnb_n3_ip_B, ngci, cpu, cpu_of_user_plane, start, duration, end, deploy_time_bias)
		}
	}
	
}