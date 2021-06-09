package main

func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	res := make([]int, n+1)
	res[1] = 1

	return sub(n, res)
}

func sub(n int, res []int) int {
	if n < 2 {
		return res[n]
	}
	if res[n] > 0 {
		return res[n]
	}

	res[n] = (sub(n-1, res) + sub(n-2, res)) % 1000000007 // 注意题意

	return res[n]
}
