package slicebinpack

// search size large first
func (n *node) find_leaf_size(width, height int, tree []*node) *node {
	max   := 0
	index := -1
	for i := 0; i < len(tree); i++ {
		if tree[i].right == nil && tree[i].top == nil && width <= tree[i].width && height <= tree[i].height {
			if max < (tree[i].width * tree[i].height) {
				max   = (tree[i].width * tree[i].height)
				index = i
			}
		}
	}
	if index != -1 {
		return tree[index]
	}else {
		return nil
	}
}