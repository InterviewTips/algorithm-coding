package _144

import "algorithm/guns"

type TreeNode = guns.TreeNode

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0) // reset
	sub(root)
	return res
}

func sub(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	sub(root.Left)
	sub(root.Right)
}
