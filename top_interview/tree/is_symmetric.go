package tree

// 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-interview-150
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTraverse(root, root)
}

func isSymmetricTraverse(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSymmetricTraverse(p.Left, q.Right) && isSymmetricTraverse(p.Right, q.Left)
}
