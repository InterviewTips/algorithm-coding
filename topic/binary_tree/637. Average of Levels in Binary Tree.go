package binary_tree

import "algorithm/template"

// 层次遍历的变种 也可以使用深度搜索
func averageOfLevels(root *template.TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*template.TreeNode, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
		var levelSum float64
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
			levelSum += float64(node.Val)
		}
		levelAverage := levelSum / float64(size)
		// 保留 .5f 官方判题样例貌似没有对此进行要求
		//value, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", levelAverage), 64)
		//res = append(res, value)
		res = append(res, levelAverage)
	}

	return res
}
