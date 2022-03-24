package _0220321

import "algorithm/topic/binary_tree"

type TreeNode = binary_tree.TreeNode

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

	data := make(map[int]int)
	for i := 0; i < len(res); i++ {
		_, ok := data[k-res[i]]
		if ok {
			return true
		}
		data[res[i]] = i
	}

	return false
}
