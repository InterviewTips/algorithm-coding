package binary_tree

import "algorithm/template"

func countNodes(root *template.TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	treeQueue := make([]*template.TreeNode, 0)
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
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
			res++
		}
	}

	return res
}
