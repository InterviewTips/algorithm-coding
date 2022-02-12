package iteration

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 左中右
	// root 没有先 push 到 stack 中 因为访问节点不是先中
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left // 左
		} else { // 走到这里说明 cur == nil, for 条件中 len(stack) != 0 一定成立
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val) // 中
			cur = cur.Right            // 右
		}
	}
	return res
}
