package binary_tree

import "algorithm/template"

type Node = template.NTreeNode

// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
// 层次遍历的变种
// 注: 方法名与原先层次遍历 levelOrder 冲突 修改为 nTreeLevelOrder
func nTreeLevelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*Node, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		levelRes := make([]int, 0)
		size := len(treeQueue)
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			// add children to queue
			treeQueue = append(treeQueue, node.Children...)
			// add node.Val to levelRes
			levelRes = append(levelRes, node.Val)
		}
		// append levelRes to res
		res = append(res, levelRes)
	}

	return res
}
