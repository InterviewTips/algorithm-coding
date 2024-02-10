package _104

import (
	"algorithm/template"
)

type TreeNode = template.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	return maxValue(left, right)
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
