package binary_tree

// 是否是对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareTree(root.Left, root.Right)
}

func compareTree(left *TreeNode, right *TreeNode) bool {
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
