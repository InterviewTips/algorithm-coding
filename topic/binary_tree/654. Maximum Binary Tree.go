package binary_tree

import "algorithm/template"

func constructMaximumBinaryTree(nums []int) *template.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := getNumsMaxIndex(nums)
	root := &template.TreeNode{Val: nums[index]}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}

func getNumsMaxIndex(nums []int) int {
	maxIndex := 0 // nums.len > 0 所以可以直接赋值 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}
