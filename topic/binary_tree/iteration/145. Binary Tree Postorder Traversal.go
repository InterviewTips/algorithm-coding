package iteration

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 { // 中右左
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 反转 左右中
	reverse(&res)
	return res
}

func reverse(res *[]int) {
	for i := 0; i < len(*res)/2; i++ {
		(*res)[i], (*res)[len(*res)-i-1] = (*res)[len(*res)-i-1], (*res)[i]
	}
}
