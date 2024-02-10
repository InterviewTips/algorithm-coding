package binary_tree

func isBalanced(root *TreeNode) bool {
	if getDepth(root) == -1 {
		return false
	}
	return true
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// 求左右子树的高度
	leftDepth := getDepth(node.Left)
	if leftDepth == -1 { // 注意要提前返回
		return -1
	}
	rightDepth := getDepth(node.Right)
	if rightDepth == -1 { // 注意要提前返回
		return -1
	}
	if getAbs(leftDepth, rightDepth) > 1 {
		return -1
	} else {
		return 1 + getMax(leftDepth, rightDepth) // 1 表示本身节点
	}
}

func getAbs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
