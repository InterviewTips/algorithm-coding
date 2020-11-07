package _768

import "sort"

func maxChunksToSorted(arr []int) int {
	count := make(map[int]int)
	ans := 0
	nozero := 0
	except := make([]int, len(arr))
	copy(except, arr)
	sort.SliceStable(except, func(i, j int) bool {
		return except[i] < except[j]
	})
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		y := except[i]

		count[x] += 1
		if count[x] == 0 {
			nozero--
		}
		if count[x] == 1 {
			nozero++
		}

		count[y] -= 1
		if count[y] == 0 {
			nozero--
		}
		if count[y] == -1 {
			nozero++
		}

		if nozero == 0 {
			ans++
		}
	}

	return ans

}
