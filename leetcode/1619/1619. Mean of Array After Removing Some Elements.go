package _619

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	for _, x := range arr[n/20 : 19*n/20] {
		sum += x
	}
	return float64(sum*10) / float64(n*9)
}
