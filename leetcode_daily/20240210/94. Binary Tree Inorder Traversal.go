package _0240210

import (
	"algorithm/template"
)

type TreeNode = template.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var sub func(n *TreeNode)
	sub = func(n *TreeNode) {
		if n == nil {
			return
		}
		sub(n.Left)
		res = append(res, n.Val)
		sub(n.Right)
	}
	sub(root)
	return res
}
