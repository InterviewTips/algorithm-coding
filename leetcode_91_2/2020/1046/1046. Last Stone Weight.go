package _046

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for len(stones) > 1 {
		n := len(stones)
		diff := stones[n-1] - stones[n-2]
		if diff > 0 {
			stones = append([]int{diff}, stones[:n-2]...)
			sort.Ints(stones)
		} else {
			stones = append([]int{}, stones[:n-2]...)
		}
	}

	if len(stones) == 0 {
		return 0
	} else {
		return stones[0]
	}
}
