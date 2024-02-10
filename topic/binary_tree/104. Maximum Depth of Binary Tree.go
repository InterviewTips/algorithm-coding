package binary_tree

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
// 可以层次遍历也可以深度遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	treeQueue := make([]*TreeNode, 0)
	// push root
	treeQueue = append(treeQueue, root)
	depth := 0
	for len(treeQueue) != 0 {
		size := len(treeQueue) // 先取出当前 queue 长度 其实就是需要将上一层的节点都取出来
		// 不能直接使用 len(treeQueue) 的原因在于 treeQueue 的长度在下面会有变动
		depth++
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			// push to queue
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
		}
	}

	return depth
}
