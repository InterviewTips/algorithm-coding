package _129

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
