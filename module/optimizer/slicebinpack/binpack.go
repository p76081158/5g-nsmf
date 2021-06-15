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

type node struct {
	x, y, width, height int
	right, top          *node
}

func (n *node) find(width, height int, algorithm string, tree []*node) *node {
	switch algorithm {
	case "right-top":
		return n.find_right_top(width, height)
	case "top-right":
		return n.find_top_right(width, height)
	case "trash-size":
		return n.find_trash_size(width, height, tree)
	case "trash-recycle":
		return n.find_top_right(width, height)
	}
	return nil

	// if n.right != nil || n.top != nil {
	// 	right := n.right.find(width, height)
	// 	if right != nil {
	// 		return right
	// 	}
	// 	return n.top.find(width, height)
	// } else if width <= n.width && height <= n.height {
	// 	return n
	// }
	// return nil
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