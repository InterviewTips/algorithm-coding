package binary_tree

func findBottomLeftValue(root *TreeNode) int {
	if root == nil { // 这种情况理论上不会发生
		return 0
	}
	var res int
	treeQueue := make([]*TreeNode, 0)
	// push root to treeQueue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
		for i := 0; i < size; i++ {
			// 每次保存第一个 最左边
			// pop_front
			node := treeQueue[0]
			if i == 0 {
				res = node.Val
			}
			treeQueue = treeQueue[1:]
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
		}
	}

	return res
}
