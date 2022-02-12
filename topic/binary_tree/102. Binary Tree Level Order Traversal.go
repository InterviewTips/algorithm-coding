package binary_tree

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
// 关键点 使用队列
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*TreeNode, 0)
	// push root
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		levelRes := make([]int, 0) // 其中一层 res
		size := len(treeQueue)     // 先取出当前 queue 长度
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			// 加入 levelRes
			levelRes = append(levelRes, node.Val)
			// push to queue
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
		}
		// levelRes append to res
		res = append(res, levelRes)
	}

	return res
}
