package slicebinpack

import (
	"fmt"
	"sort"
)

// concat top node and distribute network slice reqeust to select top nodes
func concatTop(tree_all []*node, tree_top []*node, slice Slice) ([]*node, []*node, []SliceDeploy, []DrawBlock) {
	end                := true
	deploy_list        := []SliceDeploy{}
	draw_list          := []DrawBlock{}
	tree_all_new       := []*node{}
	tree_top_candidate := []*node{}
	tree_top_new       := []*node{}
	last_width         := 0

	// fmt.Println(len(tree_all))
	// fmt.Println(len(tree_top))
	fmt.Println(slice)
	for i := 0; i < len(tree_all); i++ {
		fmt.Println(tree_all[i])
	}
	fmt.Println("")
	sort.Sort(ByNodeX(tree_top))

	for i := 0; i < len(tree_top); i++ {
		if tree_top[i].right == nil && tree_top[i].top == nil && tree_top[i].height >= slice.Height {
			tree_top_candidate = append(tree_top_candidate, tree_top[i])
		} 
	}

	for i := 0; i < len(tree_top_candidate); i++ {
		finish        := false
		tree_top_temp := []*node{}
		tree_top_temp  = append(tree_top_temp, tree_top_candidate[i])
		temp          := slice.Width - tree_top_candidate[i].width
		current_x     := tree_top_candidate[i].x + tree_top_candidate[i].width
		for j := i + 1; j < len(tree_top_candidate); j++ {
			if current_x  == tree_top_candidate[j].x {
				temp      -= tree_top_candidate[j].width
				current_x += tree_top_candidate[j].width
				tree_top_temp = append(tree_top_temp, tree_top_candidate[j])
				if temp <= 0 {
					finish     = true
					last_width = temp + tree_top_candidate[j].width
				}
			}
		}
		if finish {
			tree_top_candidate = tree_top_temp
			end                = false
			break
		}
	}


	// for i := 0; i < len(tree_top_candidate); i++ {
	// 	fmt.Println(tree_top_candidate[i])
	// 	if tree_top[i].right == nil && tree_top[i].top == nil && tree_top[i].height >= slice.Height {
	// 		for j := i + 1; j < len(tree_top); j++ {
	// 			if tree_top[j].height >= slice.Height && tree_top[i].x + tree_top[i].width == tree_top[j].x {

	// 			}
	// 		}




	// 		if temp - tree_top[i].width > 0 && slice.Height <= tree_top[i].height {
	// 			if i != len(tree_top) - 1 && (tree_top[i].x + tree_top[i].width) == tree_top[i+1].x {
	// 				temp -= tree_top[i].width
	// 				tree_index_list = append(tree_index_list, i)					
	// 			}
	// 		} else if slice.Height <= tree_top[i].height {
	// 			last_width = temp
	// 			tree_index_list = append(tree_index_list, i)
	// 			end = false
	// 			break
	// 		}
	// 	}
	// }

	if end {
		return nil, nil, nil, nil
	}

	sub_slices_num := len(slice.SubBlock)
	width_bias     := 0
	if sub_slices_num >= 1 {
		index := 0
		// node := tree_top[tree_index_list[index]]
		node := tree_top_candidate[index]
		for i := 0; i < sub_slices_num; i++ {
			var w, h int
			if node.right == nil && node.top == nil {
				if slice.SubBlock[i].Width - width_bias <= node.width {
					w = slice.SubBlock[i].Width - width_bias
					h = slice.SubBlock[i].Height

					node = node.split(w, h)
					updateTree(node, tree_top, w, h)
					tree_all_new = append(tree_all_new, node.top)
					tree_all_new = append(tree_all_new, node.right)
					tree_top_new = append(tree_top_new, node.top)
					end := false
					info := SliceDeploy {
						Name:     slice.SubBlock[i].Name,
						Ngci:     slice.Ngci,
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
					if i == sub_slices_num - 1 {
						info.End = true
					} else {
						node = node.right
					}
					width_bias = 0
					deploy_list = append(deploy_list, info)
					draw_list   = append(draw_list, drawinfo)
				} else {
					w = node.width
					h = slice.SubBlock[i].Height
					node = node.split(w, h)
					updateTree(node, tree_top, w, h)
					tree_all_new = append(tree_all_new, node.top)
					tree_all_new = append(tree_all_new, node.right)
					tree_top_new = append(tree_top_new, node.top)
					end := false
					info := SliceDeploy {
						Name:     slice.SubBlock[i].Name,
						Ngci:     slice.Ngci,
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
					width_bias = w
					index++
					i--
					// node = tree_top[tree_index_list[index]]
					node = tree_top_candidate[index]
					if i == sub_slices_num - 1 {
						info.End = true
					}
					deploy_list = append(deploy_list, info)
					draw_list   = append(draw_list, drawinfo)					
				}
			}
		}
	} else {
		for i := 0; i < len(tree_top_candidate); i++ {
			var w,h int
			// node := tree_top[tree_index_list[i]]
			node := tree_top_candidate[i]
			if i == len(tree_top_candidate) - 1 {
				w = last_width
				h = slice.Height
			} else {
				w = node.width
				h = slice.Height
			}
			node = node.split(w, h)
			updateTree(node, tree_top, w, h)
			tree_all_new = append(tree_all_new, node.top)
			tree_all_new = append(tree_all_new, node.right)
			tree_top_new = append(tree_top_new, node.top)
			end := false
			info := SliceDeploy {
				Name:     slice.Name,
				Ngci:     slice.Ngci,
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
			if i == len(tree_top_candidate) - 1 {
				info.End = true
			}
			deploy_list  = append(deploy_list, info)
			draw_list    = append(draw_list, drawinfo)
		}
	}
	return tree_all_new, tree_top_new, deploy_list, draw_list
}