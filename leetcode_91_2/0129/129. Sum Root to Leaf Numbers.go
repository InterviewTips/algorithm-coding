package _129

import (
	"github.com/InterviewTips/algorithm-coding/guns"
)

type TreeNode = guns.TreeNode

func sumNumbers(root *TreeNode) int {
	return sub(root, 0)
}

func sub(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	sum := num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return sub(root.Left, sum) + sub(root.Right, sum)
}
