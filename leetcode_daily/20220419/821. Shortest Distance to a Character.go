package _0220419

import "math"

func shortestToChar(s string, c byte) []int {
	cIndex := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cIndex = append(cIndex, i)
		}
	}

	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
		} else {
			res[i] = getMin(cIndex, i)
		}
	}

	return res
}

func getMin(cIndex []int, index int) int {
	minValue := math.MaxInt64
	for i := 0; i < len(cIndex); i++ {
		value := getAbs(cIndex[i], index)
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
