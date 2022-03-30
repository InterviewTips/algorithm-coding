package advanced

import "algorithm/topic/binary_tree"

type TreeNode = binary_tree.TreeNode

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中左右
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return res
}