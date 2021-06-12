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
	right, down         *node
}

func (n *node) find(width, height int) *node {
	if n.right != nil || n.down != nil {
		right := n.right.find(width, height)
		if right != nil {
			return right
		}
		return n.down.find(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}

func (n *node) split(width, height int) *node {
	n.down = &node{
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

func (n *node) grow(width, height int) (root, grown *node) {
	canGrowDown := width <= n.width
	canGrowRight := height <= n.height

	// attempt to keep square-ish by growing right when height is much greater than width
	shouldGrowRight := canGrowRight && (n.height >= (n.width + width))

	// attempt to keep square-ish by growing down when width is much greater than height
	shouldGrowDown := canGrowDown && (n.width >= (n.height + height))

	if shouldGrowRight {
		return n.growRight(width, height)
	} else if shouldGrowDown {
		return n.growDown(width, height)
	} else if canGrowRight {
		return n.growRight(width, height)
	} else if canGrowDown {
		return n.growDown(width, height)
	}

	// need to ensure sensible root starting size to avoid this happening
	return nil, nil
}

func (n *node) growRight(width, height int) (root, grown *node) {
	newRoot := &node{
		x:      0,
		y:      0,
		width:  n.width + width,
		height: n.height,
		down:   n,
		right: &node{
			x:      n.width,
			y:      0,
			width:  width,
			height: n.height,
		},
	}

	node := newRoot.find(width, height)
	if node != nil {
		return newRoot, node.split(width, height)
	}
	return nil, nil
}

func (n *node) growDown(width, height int) (root, grown *node) {
	newRoot := &node{
		x:      0,
		y:      0,
		width:  n.width,
		height: n.height + height,
		down: &node{
			x:      0,
			y:      n.height,
			width:  n.width,
			height: height,
		},
		right: n,
	}

	node := newRoot.find(width, height)
	if node != nil {
		return newRoot, node.split(width, height)
	}
	return nil, nil
}