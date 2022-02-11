package binary_tree

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	subPostorderTraversal(root, &res)
	return res
}

func subPostorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 左右中
	subPostorderTraversal(node.Left, res)
	subPostorderTraversal(node.Right, res)
	*res = append(*res, node.Val)

	return
}
