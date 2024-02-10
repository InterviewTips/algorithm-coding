package _4

import (
	"algorithm/template"
)

type TreeNode = template.TreeNode

func numColor(root *TreeNode) int {
	// 遍历之后进行去重
	count := 0
	data := make(map[int]struct{})
	subNumColor(root, data, &count)
	return count
}

func subNumColor(root *TreeNode, data map[int]struct{}, count *int) {
	if root == nil {
		return
	}
	_, ok := data[root.Val]
	if !ok {
		data[root.Val] = struct{}{}
		*count++
	}
	subNumColor(root.Left, data, count)
	subNumColor(root.Right, data, count)

	return
}
