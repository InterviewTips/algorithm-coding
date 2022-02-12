package iteration

// 左右中 -> 压入栈中的顺序是 (中+nil)右左
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	// push root
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil { // 添加节点到 stack 中
			stack = append(stack, node)
			stack = append(stack, nil) // 关键点
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else { // nil 需要取值到 res 结果集
			// pop 出 nil 前的元素
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, realNode.Val)
		}
	}
	return res
}
