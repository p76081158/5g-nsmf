// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binpack implements Jake Gordon's 2D binpacking algorithm.
//
// The algorithm used is described on Jake's blog here:
//
//   http://codeincomplete.com/posts/2011/5/7/bin_packing/
//
// And is also implemented by him in JavaScript here:
//
//   https://github.com/jakesgordon/bin-packing
//

package slicebinpack

// top   == left  child of tree node
// right == right child of tree node
type node struct {
	x, y, width, height int
	top, right          *node
}

func (n *node) find(width, height int, algorithm string, tree []*node) *node {
	switch algorithm {
	case "invert-pre-order":
		return n.find_invert_pre_order(width, height)
	case "pre-order":
		return n.find_pre_order(width, height)
	case "leaf-size":
		return n.find_leaf_size(width, height, tree)
	case "trash-recycle":
		return n.find_pre_order(width, height)
	}
	return nil
}

func (n *node) split(width, height int) *node {
	n.top = &node{
		x:      n.x,
		y:      n.y + height,
		width:  n.width,
		height: n.height - height,
	}

	n.right = &node{
		x:      n.x + width,
		y:      n.y,
		width:  n.width - width,
		height: n.height,
	}
	return n
}

// ByTreeX implements sort.Interface based on the node.x field.
type ByNodeX []*node

func (a ByNodeX) Len() int           { return len(a) }
func (a ByNodeX) Less(i, j int) bool { return a[i].x < a[j].x }    // ascending sort
func (a ByNodeX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func updateTree(n *node, tree []*node, width, height int) {
	for i := 0; i < len(tree); i++ {
		if tree[i].right == nil && tree[i].top == nil {
			if n.x < tree[i].x + tree[i].width && n.y > tree[i].y && n.right.x > tree[i].x && tree[i].height > n.y - tree[i].y {
				tree[i].height = n.y - tree[i].y
			}
			if n.y < tree[i].y + tree[i].height && n.x > tree[i].x && n.top.y > tree[i].y && tree[i].width > n.x - tree[i].x {
				tree[i].width = n.x - tree[i].x
			}
			// update node mapping
			if n.x == tree[i].x && n.y == tree[i].y {
				tree[i].top   = n.top
				tree[i].right = n.right
			}
		}
	}
}

func findTopRight(target *node, top []*node, right []*node) string {
	var result = "none"
	for _, t := range top {
		if t == target {
			result = "top"
		}
	}
	for _, t := range right {
		if t == target {
			result = "right"
		}
	}
	return result
}