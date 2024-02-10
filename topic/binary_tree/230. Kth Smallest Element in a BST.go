package binary_tree

func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
		return
	}

	inorder(root)

	return res[k-1]
}
