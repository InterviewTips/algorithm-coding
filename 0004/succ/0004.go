package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		c []int
		i = 0
		j = 0
	)

	c = make([]int, 0)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			c = append(c, nums1[i])
			i = i + 1
		} else {
			c = append(c, nums2[j])
			j = j + 1
		}
	}
	for i < len(nums1) {
		c = append(c, nums1[i])
		i = i + 1
	}
	for j < len(nums2) {
		c = append(c, nums2[j])
		j = j + 1
	}

	numLen := len(nums1) + len(nums2)

	if numLen%2 == 0 {
		return float64(c[numLen/2 ]+c[numLen/2-1]) / 2.0
	} else {
		return float64(c[numLen/2])
	}

}

func main() {

	//fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

}
