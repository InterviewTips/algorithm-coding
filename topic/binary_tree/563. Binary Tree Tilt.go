package binary_tree

import "algorithm/template"

func findTilt(root *template.TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	subFindTilt(root, &sum)
	return sum
}

// 返回值表示当前节点的左右子树+当前节点的值
func subFindTilt(node *template.TreeNode, sum *int) int {
	if node == nil {
		return 0
	}

	left := node.Val + subFindTilt(node.Left, sum)
	//log.Println("node", node.Val, "left is ", left)
	right := node.Val + subFindTilt(node.Right, sum)
	//log.Println("node", node.Val, "right is ", right)

	value := getAbs(left, right)
	//log.Println("node", node.Val, "value is", value)
	*sum = *sum + value

	return left + right - node.Val // 关键: left + right 计算了两遍 node.Val 所以减去一份
}
