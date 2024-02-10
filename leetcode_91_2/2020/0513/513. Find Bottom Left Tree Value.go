package _513

import (
	"algorithm/template"
)

type TreeNode = template.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	var queue []*TreeNode
	res := -1
	queue = append(queue, root)
	for len(queue) > 0 {
		res = queue[0].Val // 每一层的第一个元素的值
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return res
}
