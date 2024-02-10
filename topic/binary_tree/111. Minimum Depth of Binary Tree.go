package binary_tree

import "algorithm/template"

func minDepth(root *template.TreeNode) int {
	return subMinDepth(root)
}

func subMinDepth(root *template.TreeNode) int {
	if root == nil {
		return 0
	}

	var leftDepth int
	var rightDepth int
	if root.Left != nil { // 计算左边的
		leftDepth = 1 + subMinDepth(root.Left)
	}
	if root.Right != nil { // 计算右边的
		rightDepth = 1 + subMinDepth(root.Right)
	}

	// 等价于 root.Left == nil && root.Right == nil
	if leftDepth == 0 && rightDepth == 0 {
		return 1 // 这里是 1
	}
	if leftDepth == 0 {
		return rightDepth
	}
	if rightDepth == 0 {
		return leftDepth
	}

	// 都不为 0
	if leftDepth < rightDepth {
		return leftDepth
	}

	return rightDepth
}

// 使用层次遍历来做
// 关键点: 左右孩子都为空
// 提交时需要改函数名 minDepthLevel -> minDepth
func minDepthLevel(root *template.TreeNode) int {
	if root == nil {
		return 0
	}
	treeQueue := make([]*template.TreeNode, 0)
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
			if node.Left == nil && node.Right == nil { // 关键
				return depth
			}
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
