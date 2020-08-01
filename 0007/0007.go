package main

import "fmt"

// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
// 请根据这个假设，如果反转后整数溢出那么就返回 0。

const INT_MIN = -(1 << 31)
const INT_MAX = 1<<31 - 1

func reverse(x int) int {
	// 10 进制反转
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp < INT_MIN || tmp > INT_MAX {
		tmp = 0
	}

	return tmp
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}
