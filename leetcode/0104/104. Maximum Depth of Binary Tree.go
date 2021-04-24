package _104

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func maxDepth(root *TreeNode) int {
	return subMaxDepth(root)
}

func subMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+subMaxDepth(root.Left), 1+subMaxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
