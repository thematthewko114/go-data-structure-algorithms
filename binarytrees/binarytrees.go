package binarytrees

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.InsertNode(data)
	}
	return t
}

func (n *BinaryNode) InsertNode(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.InsertNode(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.InsertNode(data)
		}
	}
}

func (t *BinaryTree) GetRoot() *BinaryNode {
	return t.root
}

func (n *BinaryNode) PrintTree(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, n.data)
	n.left.PrintTree(w, ns+2, 'L')
	n.right.PrintTree(w, ns+2, 'R')
}
