package tree

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	root := &TreeNode{Val: 1}
	right := &TreeNode{Val: 2}
	root.Right = right
	ret := maxDepth(root)
	fmt.Println(ret)
}

func Test_isSameTree(t *testing.T) {
	p := &TreeNode{Val: 1}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	ret := isSameTree(p, q)
	fmt.Println("ret", ret)
}

func Test_isSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}
	ret := isSymmetric(root)
	fmt.Println("ret", ret)
}
