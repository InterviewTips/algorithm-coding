package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var (
		b, i, j int
		ok      bool
		m       map[int]int
	)

	m = make(map[int]int, len(nums))
	for i, b = range nums {
		if j, ok = m[target-b]; ok {
			// 在 i 之前存在 j，使得 a + b = target
			return []int{j, i}
		}

		// 存进 map 中，通过 m[整数] 可以获取到对应索引
		m[b] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{7, 11, 2, 15}, 9))
}
