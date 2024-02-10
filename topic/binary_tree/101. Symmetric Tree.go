package binary_tree

import "algorithm/template"

// 是否是对称的二叉树
func isSymmetric(root *template.TreeNode) bool {
	if root == nil {
		return true
	}
	return compareTree(root.Left, root.Right)
}

func compareTree(left *template.TreeNode, right *template.TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	if left.Val != right.Val {
		return false
	}
	outside := compareTree(left.Left, right.Right)
	inside := compareTree(left.Right, right.Left)

	return outside && inside
}
