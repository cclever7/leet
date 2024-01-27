package tree

// 相同的树
// https://leetcode.cn/problems/same-tree/description/?envType=study-plan-v2&envId=top-interview-150
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return traverseIsSameTree(p, q)
}

func traverseIsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return traverseIsSameTree(p.Left, q.Left) && traverseIsSameTree(p.Right, q.Right)
}
