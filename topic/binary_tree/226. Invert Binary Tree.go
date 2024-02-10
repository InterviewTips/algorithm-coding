package binary_tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 交换左右节点
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
