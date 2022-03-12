package leetcode

func findSecondMinimumValue(root *TreeNode) int {

	res := -1
	rootVal := root.Val
	var myFunc func(node *TreeNode)
	myFunc = func(node *TreeNode) {
		if node == nil || (res != -1 && node.Val >= res) { // node.Val >= res 说明不需要进行更新 res
			return
		}

		if node.Val > rootVal {
			res = node.Val
		}

		myFunc(node.Left)
		myFunc(node.Right)
	}

	myFunc(root)
	// fmt.Println(res)

	return res

}
