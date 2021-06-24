package slicebinpack

// tree traversal: pre-order traversal
func (n *node) find_pre_order(width, height int) *node {
	if n.right != nil || n.top != nil {
		top := n.top.find_pre_order(width, height)
		if top != nil {
			return top
		}
		return n.right.find_pre_order(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}