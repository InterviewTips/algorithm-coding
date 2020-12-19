package _814

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}
