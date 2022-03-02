package theory

import "log"

// 多重背包可以化解为 01 背包
func multiplePack(weight, value, nums []int, bagWeight int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = getMax(res[j], res[j-weight[i]]+value[i])
		}
		log.Println(res)
	}

	return res[bagWeight]
}
