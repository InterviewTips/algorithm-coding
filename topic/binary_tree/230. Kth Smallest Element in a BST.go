package binary_tree

import "algorithm/template"

func kthSmallest(root *template.TreeNode, k int) int {
	res := make([]int, 0)
	var inorder func(node *template.TreeNode)
	inorder = func(node *template.TreeNode) {
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
