package dp

import "log"

func countSubstrings(s string) int {
	res := make([]int, len(s))

	for i := 0; i < len(res); i++ {
		res[i] = 1
	}
	count := res[0]
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if isValid(s[j : i+1]) {
				res[i] += 1
			}
		}
		log.Println(res[i])
		count += res[i]
	}
	log.Println(res)

	return count
}

func isValid(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
