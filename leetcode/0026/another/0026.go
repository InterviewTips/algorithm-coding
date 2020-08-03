package main

import "fmt"

func removeDuplicates(nums []int) int {
	var (
		nLen = len(nums)
		i    int
	)

	for i = 0; i < nLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			j := i
			for j < nLen-1 {
				nums[j] = nums[j+1]
				j++
			}
			nLen = nLen - 1
			i--
		}
	}

	fmt.Println(nums)

	return nLen
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
