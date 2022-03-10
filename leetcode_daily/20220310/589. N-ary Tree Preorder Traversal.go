package _0220310

import "github.com/InterviewTips/algorithm-coding/topic/binary_tree"

type Node = binary_tree.Node

func preorder(root *Node) []int {
	res := make([]int, 0)
	var subPreorder func(node *Node)
	subPreorder = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		for i := 0; i < len(node.Children); i++ {
			subPreorder(node.Children[i])
		}
	}

	subPreorder(root)
	return res
}
