package main

import "fmt"

// 斐波那契数列
// f(n) = 1, n = 1
// f(n) = 2, n = 2
// f(n) = f(n-1) + f(n-2), n >= 3

//func climbStairs(n int) int {
//	// n 为正整数
//	if n == 1 || n == 2 {
//		return n
//	}
//
//	return climbStairs(n-1) + climbStairs(n-2)
//}

func climbStairs(n int) int {
	// 非递归写法
	if n == 1 || n == 2 {
		return n
	}
	dp := [3]int{1, 2}
	for i := 2; i < n; i++ {
		dp[2] = dp[0] + dp[1]
		dp[0] = dp[1] // 替换
		dp[1] = dp[2] // 替换

		// 也可以写成这样 容易理解
		//t := dp[1]
		//dp[1] = dp[2]
		//dp[0] = t
	}
	return dp[2]
}

func main() {
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(3))
}
