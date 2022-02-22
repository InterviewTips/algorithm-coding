package greedy

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0           // 计数
	res := math.MinInt64 // 最终结果
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res { // 要先变动先 这样 math.MinInt64 才会被修改
			res = count
		}
		if count <= 0 {
			count = 0 // 重新计数
		}
	}

	return res
}

// 暴力破解超时
func ErrMaxSubArray(nums []int) int {
	res := math.MinInt64
	for i := 1; i <= len(nums); i++ {
		// 滑动窗口的大小
		for index := 0; index < len(nums)-i+1; index++ {
			if value := getSum(nums[index : index+i]); value > res {
				res = value
			}
		}
	}
	return res
}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
