package leetcode

import (
	"algorithm/template"
)

type TreeNode = template.TreeNode

func findTarget(root *TreeNode, k int) bool {
	res := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	// 两数之和
	data := make(map[int]struct{})
	target := k
	for i := 0; i < len(res); i++ {
		_, ok := data[target-res[i]]
		if ok {
			return true
		}
		data[res[i]] = struct{}{}
	}

	return false
}
