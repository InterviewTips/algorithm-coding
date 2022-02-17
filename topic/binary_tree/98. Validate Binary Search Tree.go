package binary_tree

// 二叉搜索树 每个节点的值都不能重复
func isValidBST(root *TreeNode) bool {
	// 利用性质 中序遍历有序
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	// push root to stack
	stack = append(stack, root)
	for len(stack) != 0 {
		// pop node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else { // nil pop realNode
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(res) > 0 && res[len(res)-1] >= realNode.Val { // 直接 return
				return false
			} else { // 加入 res 中
				res = append(res, realNode.Val)
			}
		}
	}

	return true
}
