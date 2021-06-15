package slicebinpack

// right first search
func (n *node) find_right_top(width, height int) *node {
	if n.right != nil || n.top != nil {
		right := n.right.find_right_top(width, height)
		if right != nil {
			return right
		}
		return n.top.find_right_top(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}