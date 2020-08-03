package main

import "fmt"

func removeElement(nums []int, val int) int {
	var (
		i int
		j = 0
	)

	for i = 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j

}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
