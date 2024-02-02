package _686

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	totalValues := make([][2]int, n)

	for i := 0; i < n; i++ {
		totalValues[i] = [2]int{i, aliceValues[i] + bobValues[i]}
	}

	sort.Slice(totalValues, func(i, j int) bool {
		return totalValues[i][1] > totalValues[j][1]
	})

	aliceScore, bobScore := 0, 0
	for i, pair := range totalValues {
		if i%2 == 0 {
			aliceScore += aliceValues[pair[0]]
		} else {
			bobScore += bobValues[pair[0]]
		}
	}

	if aliceScore > bobScore {
		return 1
	} else if aliceScore < bobScore {
		return -1
	} else {
		return 0
	}
}
