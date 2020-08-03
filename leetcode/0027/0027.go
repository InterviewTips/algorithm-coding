package main

import "fmt"

func removeElement(nums []int, val int) int {
	var (
		nLen = len(nums)
		i    int
	)

	for i = 0; i < nLen; i++ {
		if nums[i] == val {
			j := i
			for j < nLen-1 {
				nums[j] = nums[j+1]
				j++
			}
			nLen = nLen - 1
			i--
		}
	}

	//fmt.Println(nums)

	return nLen

}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
