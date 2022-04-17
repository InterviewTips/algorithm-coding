package _022season

import "math"

func giveGem(gem []int, operations [][]int) int {
	// 注意是一半
	for i := 0; i < len(operations); i++ {
		item := operations[i]
		value := gem[item[0]] / 2
		gem[item[0]] -= value
		gem[item[1]] += value
	}

	minValue := math.MaxInt64
	maxValue := math.MinInt64
	for i := 0; i < len(gem); i++ {
		if gem[i] > maxValue {
			maxValue = gem[i]
		}
		if gem[i] < minValue {
			minValue = gem[i]
		}
	}

	return maxValue - minValue
}
