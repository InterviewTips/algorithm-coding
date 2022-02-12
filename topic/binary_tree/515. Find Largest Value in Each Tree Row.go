package binary_tree

import (
	"math"
)

// 给定一棵二叉树的根节点 root, 请找出该二叉树中每一层的最大值
// 层次遍历的变种
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*TreeNode, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
		levelMax := math.MinInt64
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
			if node.Val > levelMax {
				levelMax = node.Val
			}
		}
		// push levelMax to res
		res = append(res, levelMax)
	}

	return res
}
