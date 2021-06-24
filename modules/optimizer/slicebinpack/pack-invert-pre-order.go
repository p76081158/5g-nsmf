package slicebinpack

// tree traversal: invert binary tree and pre-order traversal (right first search in origin binary tree)
func (n *node) find_invert_pre_order(width, height int) *node {
	if n.right != nil || n.top != nil {
		right := n.right.find_invert_pre_order(width, height)
		if right != nil {
			return right
		}
		return n.top.find_invert_pre_order(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}