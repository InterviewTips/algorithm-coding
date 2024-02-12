package _0240212

import "algorithm/template"

type TreeNode = template.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var sub func(n *TreeNode)
	sub = func(n *TreeNode) {
		if n == nil {
			return
		}
		sub(n.Left)
		sub(n.Right)
		res = append(res, n.Val)
	}
	sub(root)
	return res
}
