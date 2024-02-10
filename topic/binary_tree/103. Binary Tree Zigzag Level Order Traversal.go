package binary_tree

import "algorithm/template"

// 层次遍历变种
func zigzagLevelOrder(root *template.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*template.TreeNode, 0)
	// push root
	treeQueue = append(treeQueue, root)
	positive := true
	for len(treeQueue) != 0 {
		levelRes := make([]int, 0) // 其中一层 res
		size := len(treeQueue)     // 先取出当前 queue 长度 其实就是需要将上一层的节点都取出来
		// 不能直接使用 len(treeQueue) 的原因在于 treeQueue 的长度在下面会有变动
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
		if !positive { // 逆向
			reverseLevelRes(&levelRes)
		}
		positive = !positive
		// levelRes append to res
		res = append(res, levelRes)
	}

	return res
}

func reverseLevelRes(res *[]int) {
	for i := 0; i < len(*res)/2; i++ {
		(*res)[i], (*res)[len(*res)-i-1] = (*res)[len(*res)-i-1], (*res)[i]
	}
}
