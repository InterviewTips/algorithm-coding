package _0240211

import "algorithm/template"

type TreeNode = template.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var sub func(n *TreeNode)
	sub = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		sub(n.Left)
		sub(n.Right)
	}
	sub(root)
	return res
}
