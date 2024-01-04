package treeAlgo

import (
	"fmt"
	"io"
)

type TreeNode struct {
	data   string
	weight int
	left   *TreeNode
	right  *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (t *BinaryTree) Insert(data string, weight int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{data, weight, nil, nil}
	} else {
		t.root.InsertNode(data, weight)
	}
	return t
}

func (n *TreeNode) InsertNode(data string, weight int) {
	if n == nil {
		return
	} else if weight <= n.weight {
		if n.left == nil {
			n.left = &TreeNode{data, weight, nil, nil}
		} else {
			n.left.InsertNode(data, weight)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data, weight, nil, nil}
		} else {
			n.right.InsertNode(data, weight)
		}
	}
}

func CreateRootNode(value string, weight int) *TreeNode {
	return &TreeNode{value, weight, nil, nil}
}

func PreorderTraversal(root *TreeNode) []string {
	output := make([]string, 0)
	if root == nil {
		return output
	}
	left := PreorderTraversal(root.left)
	right := PreorderTraversal(root.right)
	output = append(output, root.data)
	output = append(output, left...)
	output = append(output, right...)
	return output
}

func InorderTraversal(root *TreeNode) []string {
	output := make([]string, 0)
	if root == nil {
		return output
	}
	output = append(output, InorderTraversal(root.left)...)
	output = append(output, root.data)
	output = append(output, InorderTraversal(root.right)...)
	return output
}

func PostorderTraversal(root *TreeNode) []string {
	output := make([]string, 0)
	if root == nil {
		return output
	}
	output = append(output, PostorderTraversal(root.left)...)
	output = append(output, PostorderTraversal(root.right)...)
	output = append(output, root.data)
	return output
}

func (n *TreeNode) PrintTree(w io.Writer, ns int, ch rune) {
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

func (t *BinaryTree) GetRoot() *TreeNode {
	return t.root
}
