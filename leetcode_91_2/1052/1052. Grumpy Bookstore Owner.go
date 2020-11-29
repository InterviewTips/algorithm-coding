package _052

import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, X int) int {
	//先把所有算出来
	res := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}
	max := res
	fmt.Printf("prev res is %v\n", res)
	for i := 0; i < len(customers); i++ {
		if i < X {
			if grumpy[i] == 1 {
				res += customers[i]
			}
		} else {
			if grumpy[i-X] == 1 {
				res -= customers[i-X]
			}
			if grumpy[i] == 1 {
				res += customers[i]
			}
		}
		if res > max {
			max = res
		}
	}
	return max
}
