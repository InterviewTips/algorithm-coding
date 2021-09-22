package array

import (
	"log"
	"math"
	"sort"
)

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32 // 先设置为最大
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			subLength := right - left + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[left] // 重点 左指针
			left++            // 滑动
		}
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result
}

func minSubArrayLen1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1] // i-1 而不是 i 这样可以从 0 开始遍历
		bound := sort.SearchInts(sums, target)
		log.Println("i bound, target", i, bound, target)
		//if bound < 0 {
		//	bound = -bound - 1
		//}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
