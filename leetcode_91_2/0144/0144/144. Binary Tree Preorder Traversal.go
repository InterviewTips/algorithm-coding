package _144

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil { // 遍历左分支
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1] // 出栈
	}
	return res
}
