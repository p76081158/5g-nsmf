package slicebinpack

import (
	"fmt"
)

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

type DrawBlock struct {
	TopLeftX   int
	TopLeftY   int
	DownRightX int
	DownRightY int
}

type Packer struct {
	Bins         Bin
	AcceptSlices []Slice
	RejectSlices []Slice
	DeployInfos  []SliceDeploy
	DrawInfos    []DrawBlock
}

type Bin struct {
	Name   string
	Width  int
	Height int
	Slices []Slice
}

func (p *Packer) Pack(algorithm string) {
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
	tree_list       := []*node{ &root }
	tree_list_right := []*node{}
	tree_list_top   := []*node{}
	for i := 0; i < slices_num; i++ {
		var w, h int
		sub_slices_num := len(p.Bins.Slices[i].SubBlock)

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

		// select algorithm for finding candidate of network slice placement
		node := root.find(w, h, algorithm, tree_list)
		fmt.Println(findTopRight(node, tree_list_right, tree_list_top))
		if node != nil {
			// is slice group or not
			if sub_slices_num > 1 {
				for j :=0; j < sub_slices_num; j++ {
					w = p.Bins.Slices[i].SubBlock[j].Width
                    h = p.Bins.Slices[i].SubBlock[j].Height
					node = node.split(w, h)
					tree_list       = append(tree_list, node.right)
					tree_list       = append(tree_list, node.top)
					tree_list_right = append(tree_list_right, node.right)
					tree_list_top   = append(tree_list_top, node.top)
					end := false
					info := SliceDeploy {
						Name:     p.Bins.Slices[i].SubBlock[j].Name,
						Ngci:     p.Bins.Slices[i].Ngci,
						Start:    node.x,
						Duration: w,
						End:      end,
						Resource: h,
					}
					drawinfo := DrawBlock {
						TopLeftX:   node.x,
						TopLeftY:   node.y + h,
						DownRightX: node.x + w,
						DownRightY: node.y,
					}
					if j == sub_slices_num - 1 {
						info.End = true
					} else {
						node = node.right
					}
					p.DeployInfos = append(p.DeployInfos, info)
					p.DrawInfos   = append(p.DrawInfos, drawinfo)
				}
				p.AcceptSlices = append(p.AcceptSlices, p.Bins.Slices[i])
			} else {
				node = node.split(w, h)
				tree_list = append(tree_list, node.right)
				tree_list = append(tree_list, node.top)
				tree_list_right = append(tree_list_right, node.right)
				tree_list_top   = append(tree_list_top, node.top)
				info := SliceDeploy {
					Name:     p.Bins.Slices[i].Name,
					Ngci:     p.Bins.Slices[i].Ngci,
					Start:    node.x,
					Duration: w,
					End:      true,
					Resource: h,
				}
				drawinfo := DrawBlock {
					TopLeftX:   node.x,
					TopLeftY:   node.y + h,
					DownRightX: node.x + w,
					DownRightY: node.y,
				}
			    p.AcceptSlices = append(p.AcceptSlices, p.Bins.Slices[i])
				p.DeployInfos  = append(p.DeployInfos, info)
				p.DrawInfos    = append(p.DrawInfos, drawinfo)
			}
		} else {
			p.RejectSlices = append(p.RejectSlices, p.Bins.Slices[i])
		}
	}
}