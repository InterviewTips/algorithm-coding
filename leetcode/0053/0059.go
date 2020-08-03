package main

//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//maxSubArray
//https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1] // 改变 nums[i],后面的 nums[i+1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
