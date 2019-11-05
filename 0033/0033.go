package main

import "fmt"

func search(nums []int, target int) int {
	var (
		b, i, j int
		ok      bool
		m       map[int]int
	)

	m = make(map[int]int, len(nums))
	for i, b = range nums {
		m[b] = i
	}

	if j, ok = m[target]; ok {
		return j
	} else {
		return -1
	}

}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2},3))
}
