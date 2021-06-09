package main

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	res := make([]int, n+1)
	return sub(n, res)
}

func sub(n int, res []int) int {
	if n == 1 || n == 2 {
		res[n] = n
		return res[n]
	}
	if res[n] > 0 {
		return res[n]
	}

	res[n] = (sub(n-1, res) + sub(n-2, res)) % 1000000007

	return res[n]
}
