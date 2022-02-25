package greedy

func minCameraCover(root *TreeNode) int {

	res := 0
	// 0 无覆盖
	// 1 有摄像头
	// 2 有覆盖
	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		if node == nil { // 空节点表示有覆盖
			return 2
		}

		left := postOrder(node.Left)
		right := postOrder(node.Right)

		if left == 2 && right == 2 { // 左右节点已经都有覆盖
			return 0 // 父节点先无覆盖
		}

		// 有先后顺序
		// q1
		if left == 0 || right == 0 { // 有一个没有被覆盖
			res++
			return 1 // 父节点安装摄像头
		}

		// q2
		if left == 1 || right == 1 { // 有一个安装了摄像头
			return 2 // 父节点覆盖
		}

		return -1
	}

	if postOrder(root) == 0 {
		res++
	}
	return res
}
