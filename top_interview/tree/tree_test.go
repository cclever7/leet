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
