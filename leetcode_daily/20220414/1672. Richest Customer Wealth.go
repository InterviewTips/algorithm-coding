package _0220414

func maximumWealth(accounts [][]int) int {
	maxValue := 0
	for i := 0; i < len(accounts); i++ {
		value := getSum(accounts[i])
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue

}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
