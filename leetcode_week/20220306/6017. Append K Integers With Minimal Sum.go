package _0220306

import (
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	//log.Println(nums)
	res := make([]int, 0)
	slowIndex := 0
	res = append(res, nums[slowIndex])
	for i := 1; i < len(nums); i++ {
		if res[slowIndex] != nums[i] {
			res = append(res, nums[i])
			slowIndex++
		}
	}

	nums = res
	//log.Println(nums)

	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}
		if i == 0 {
			value := 1
			for value < nums[i] && k != 0 {
				//log.Println("add value", value)
				sum += value
				value++
				k--
			}
			//log.Println("k is", k)
			if k == 0 {
				goto end
			}
		} else {
			value := nums[i-1] + 1
			for value < nums[i] && k != 0 {
				//log.Println("i > 0, add value", value)
				sum += value
				value++
				k--
			}
			//log.Println("i > 0, k is", k)
			if k == 0 {
				goto end
			}
		}
	}
end:
	// 兜底
	value := nums[len(nums)-1]
	newValue := (2*value + 1 + k) * k / 2
	sum += newValue
	//for k != 0 {
	//	sum += value
	//	value++
	//	k--
	//}

	return int64(sum)
}
