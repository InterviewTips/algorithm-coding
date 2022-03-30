package advanced

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左右中
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}

	postorder(root)
	return res
}
