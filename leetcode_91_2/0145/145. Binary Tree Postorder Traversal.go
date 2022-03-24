package _145

import "algorithm/guns"

type TreeNode = guns.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
	return res
}
