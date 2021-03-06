package _094

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}
