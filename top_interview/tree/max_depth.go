package tree

var max = 0

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func maxDepth(root *TreeNode) int {
	max = 0
	if root == nil {
		return max
	}
	traverse(root, max)
	return max
}

func traverse(root *TreeNode, depth int) {
	depth = depth + 1
	if depth > max {
		max = depth
	}

	if root.Left != nil {
		traverse(root.Left, depth)
	}
	if root.Right != nil {
		traverse(root.Right, depth)
	}
}
