package dp

import "log"

func findTargetSumWays(nums []int, target int) int {
	// leftSum - rightSum = target
	// leftSum + rightSum = sum
	// leftSum - (sum - leftSum) = target
	// leftSum = (target + sum) / 2 注: 这里 target + sum 不需要考虑溢出问题 题目给了范围 不会溢出的
	// res[j] += res[j-nums[i]]
	// res[j] 表示凑齐容量 j 有几种方式

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if getAbs(target) > sum {
		return 0
	}

	if (target+sum)&1 == 1 { // 奇数也不符合需求
		return 0
	}

	leftSum := (target + sum) / 2
	res := make([]int, leftSum+1)
	res[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := leftSum; j >= nums[i]; j-- {
			res[j] += res[j-nums[i]]
		}
		log.Println(res)
	}

	//log.Println(res)

	return res[leftSum]
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
