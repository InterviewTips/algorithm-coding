package _0220408

import "algorithm/topic/binary_tree"

type Node = binary_tree.NTreeNode

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*Node{root}
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		res = append(res, level)
	}

	return res
}
