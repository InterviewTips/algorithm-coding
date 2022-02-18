package binary_tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 从中间开始构建
	// 左开右闭
	left := 0
	// right := len(nums) // 都可以 右开
	right := len(nums) - 1 // 右闭
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
