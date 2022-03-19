package leetcode_interview

import "sort"

func CheckPermutation(s1 string, s2 string) bool {
	s1B := []byte(s1)
	s2B := []byte(s2)
	sort.SliceStable(s1B, func(i, j int) bool {
		return s1B[i] < s1B[j]
	})
	sort.SliceStable(s2B, func(i, j int) bool {
		return s2B[i] < s2B[j]
	})

	return string(s1B) == string(s2B)
}
