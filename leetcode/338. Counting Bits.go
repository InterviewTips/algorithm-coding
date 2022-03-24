package leetcode

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = Count(i)
	}

	return res
}

func Count(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}
