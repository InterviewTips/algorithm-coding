package binary_tree

// 层次遍历的变种 rightRes 关键点
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*TreeNode, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		var rightRes int
		size := len(treeQueue)
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			if i == size-1 { // 当前层的最右边 node
				rightRes = node.Val
			}
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
		}
		res = append(res, rightRes)
	}

	return res
}
