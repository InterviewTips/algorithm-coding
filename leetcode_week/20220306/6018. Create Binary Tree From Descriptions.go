package _0220306

import (
	"github.com/InterviewTips/algorithm-coding/topic/binary_tree"
)

type TreeNode = binary_tree.TreeNode

func createBinaryTree(descriptions [][]int) *TreeNode {
	data := make(map[int]*TreeNode)
	// 如何确定 root 节点
	childData := make(map[int]struct{})
	var rootVal int
	for i := 0; i < len(descriptions); i++ {
		child := descriptions[i][1]
		var node *TreeNode
		preNode, ok := data[descriptions[i][0]]
		if !ok {
			node = &TreeNode{Val: descriptions[i][0]}
		} else {
			node = preNode
		}
		childNode, ok := data[child]
		//log.Println("child ", child, "ok:", ok)
		if descriptions[i][2] == 1 {
			if ok {
				node.Left = childNode
			} else {
				node.Left = &TreeNode{Val: child}
				data[child] = node.Left
			}
		} else {
			if ok {
				node.Right = childNode
			} else {
				node.Right = &TreeNode{Val: child}
				data[child] = node.Right
			}
		}
		data[node.Val] = node
		childData[child] = struct{}{}
	}

	for i := 0; i < len(descriptions); i++ {
		nodeVal := descriptions[i][0]
		_, ok := childData[nodeVal]
		if !ok {
			rootVal = nodeVal
			break
		}
	}

	return data[rootVal]
}
