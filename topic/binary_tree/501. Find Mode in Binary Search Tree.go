package binary_tree

import "sort"

// 中序遍历 + map 空间貌似用的多
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	data := make(map[int]int)
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			// 左中右 -> 右(中+nil)左
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else { // nil readNode
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			data[realNode.Val]++
		}
	}

	// make slice
	dataArray := make([][2]int, 0)
	maxCount := -1
	for k, v := range data {
		dataArray = append(dataArray, [2]int{k, v})
		if v > maxCount {
			maxCount = v
		}
	}

	// 从大到小排列
	sort.SliceStable(dataArray, func(left, right int) bool {
		return dataArray[left][1] > dataArray[right][1]
	})

	res := make([]int, 0)
	for i := 0; i < len(dataArray); i++ {
		if dataArray[i][1] == maxCount {
			res = append(res, dataArray[i][0])
		} else { // 后面的都是 < maxCount 的
			break
		}
	}

	return res
}
