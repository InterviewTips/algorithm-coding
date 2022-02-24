package greedy

// 暴力解法
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		index := i
		rest := gas[index] - cost[index]
		index = (index + 1) % len(gas)
		for rest >= 0 && index != i {
			rest = rest + gas[index] - cost[index]
			index = (index + 1) % len(gas)
		}
		if rest >= 0 && index == i {
			return index
		}
	}

	return -1
}

// var greedyCanCompleteCircuit = canCompleteCircuit
func greedyCanCompleteCircuit(gas []int, cost []int) int {
	start := 0
	curSum := 0
	totalSum := 0
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i] // 此加油站和消耗对应剩余油量
		curSum += rest
		totalSum += rest
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}

	return start
}