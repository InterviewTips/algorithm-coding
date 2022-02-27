package dp

func lastStoneWeightII(stones []int) int {

	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	bagWeight := sum / 2

	res := make([]int, bagWeight+1)
	res[0] = 0
	for i := 0; i < len(stones); i++ {
		for j := bagWeight; j >= stones[i]; j-- {
			res[j] = getMax(res[j], res[j-stones[i]]+stones[i])
		}
	}

	bagSum := res[bagWeight]
	left := sum - bagSum
	if left > bagSum {
		return left - bagSum
	}
	return bagSum - left

}
