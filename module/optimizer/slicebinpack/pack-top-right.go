package slicebinpack

// top first search
func (n *node) find_top_right(width, height int) *node {
	if n.right != nil || n.top != nil {
		top := n.top.find_top_right(width, height)
		if top != nil {
			return top
		}
		return n.right.find_top_right(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}