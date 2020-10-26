package _144

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func preorderTraversal(root *TreeNode) []int {
	return sub(root)
}

func sub(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {
		res = append(res, sub(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, sub(root.Right)...)
	}
	return res
}
