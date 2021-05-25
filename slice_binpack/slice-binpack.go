package slice_binpack

type Block struct {
	Name     string
	Width    int
	Height   int
}

type Slice struct {
	Name     string
	Width    int
	Height   int
	Ngci     string
	SubBlock []Block
}

type SliceDeploy struct {
	Name     string
	Ngci     string
	Start    int
	Duration int
	End      bool
	Resource int
}

type Packer struct {
	Bins         Bin
	AcceptSlices []Slice
	RejectSlices []Slice
	DeployInfos  []SliceDeploy
}

type Bin struct {
	Name   string
	Width  int
	Height int

	Slices []Slice
}

func (p *Packer) Pack() {
	root := node {
		x:      0,
		y:      0,
		width:  p.Bins.Width,
		height: p.Bins.Height,
	}
	slices_num := len(p.Bins.Slices)
	if slices_num == 0 {
		return
	}
	for i := 0; i < slices_num; i++ {
		var w, h int
		sub_slices_num := len(p.Bins.Slices[i].SubBlock)
		// for case that only change resource used and not split to multiple blocks
		// if sub_slices_num == 1 {
		// 	w = p.Bins.Slices[i].SubBlock[0].Width
		// 	h = p.Bins.Slices[i].SubBlock[0].Height
		// max_h := 0
		// max_w := 0

		// choose which node for putting the slice (could be multiple blocks)
		w, h = 0, 0
		if sub_slices_num >= 1 {
			for j := 0; j < sub_slices_num; j++ {
				if p.Bins.Slices[i].SubBlock[j].Height > h {
					h = p.Bins.Slices[i].SubBlock[j].Height
				}
				w += p.Bins.Slices[i].SubBlock[j].Width
			}
		} else {
			w = p.Bins.Slices[i].Width
			h = p.Bins.Slices[i].Height
		}

		node := root.find(w, h)
		// for finding the node fit the ma
		// if sub_slices_num > 1 {
		// 	for j := 0; j < sub_slices_num; j++ {
		// 		if p.Bins.Slices[i].SubBlock[j].Height > max_h {
		// 			max_h = p.Bins.Slices[i].SubBlock[j].Height
		// 		}
		// 		max_w += p.Bins.Slices[i].SubBlock[j].Width
		// 	}
		// 	node = root.find(max_w, max_h)
		// }
		if node != nil {
			if sub_slices_num > 1 {
				for j :=0; j < sub_slices_num; j++ {
					w = p.Bins.Slices[i].SubBlock[j].Width
                    h = p.Bins.Slices[i].SubBlock[j].Height
					node = node.split(w, h)
					end := false
					info := SliceDeploy {
						Name:     p.Bins.Slices[i].SubBlock[j].Name,
						Ngci:     p.Bins.Slices[i].Ngci,
						Start:    node.x,
						Duration: w,
						End:      end,
						Resource: h,
					}
					if j == sub_slices_num - 1 {
						info.End = true
					} else {
						node = node.right
					}
					p.DeployInfos  = append(p.DeployInfos, info)
				}
				p.AcceptSlices = append(p.AcceptSlices, p.Bins.Slices[i])
			} else {
				node = node.split(w, h)
				info := SliceDeploy {
					Name:     p.Bins.Slices[i].Name,
					Ngci:     p.Bins.Slices[i].Ngci,
					Start:    node.x,
					Duration: w,
					End:      true,
					Resource: h,
				}
			    p.AcceptSlices = append(p.AcceptSlices, p.Bins.Slices[i])
				p.DeployInfos  = append(p.DeployInfos, info)
			}
		} else {
			p.RejectSlices = append(p.RejectSlices, p.Bins.Slices[i])
		}
	}
}

