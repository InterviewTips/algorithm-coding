package binary_tree

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	subPreorderTraversal(root, &res)
	return res
}

func subPreorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 中左右
	*res = append(*res, node.Val)
	subPreorderTraversal(node.Left, res)
	subPreorderTraversal(node.Right, res)

	return
}
