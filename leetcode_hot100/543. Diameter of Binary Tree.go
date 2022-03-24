package leetcode_hot100

import "algorithm/topic/binary_tree"

type TreeNode = binary_tree.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {

	res := 0
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := getDepth(node.Left)
		right := getDepth(node.Right)

		// 经过此节点的路径 例如一个深度是 3(说明路径是2)，一个深度是 2(说明路径是1)，再经过根节点+2，相互抵消 -> left + right
		res = getMax(res, left+right)

		depth := getMax(left, right) + 1
		return depth
	}

	getDepth(root)

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
