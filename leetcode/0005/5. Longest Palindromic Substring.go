package _005

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 { // i == j
				dp[i][j] = 1
			} else if l == 1 { // j i 紧跟 长度为 2
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else { // 长度大于 2
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
		//只会用到右上角部分 i + l
		fmt.Printf("dp is %v\n", dp)
	}
	return ans
}
