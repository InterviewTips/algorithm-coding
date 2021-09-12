package array

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：
//
//    你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func searchRange(nums []int, target int) []int {
	// 结合 35 题，先找出第一个大于等于 target 的索引作为 leftIndex
	// rightIndex 则是利用数组单调递增属性找到第一个大于等于 target+1 的索引减去 1
	// 为什么是可以 target+1 因为 target+1 则是第一个大于 target 的数字
	leftIndex := searchInsert(nums, target)
	rightIndex := searchInsert(nums, target+1)
	//log.Println("left, right index is", leftIndex, rightIndex)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	// 这里为什么不需要再判断 rightIndex 因为 rightIndex-1 之后必定会等于 target 的 index
	// 可能是第一个出现的 target(假如 target 只有一个的话)
	return []int{leftIndex, rightIndex - 1}

}
