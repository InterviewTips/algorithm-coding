package string

import "strings"

func isPalindrome(s string) bool {
	newS := make([]byte, 0)
	s = strings.ToUpper(s)
	for i := 0; i < len(s); i++ {
		if ('0' <= s[i] && s[i] <= '9') ||
			('A' <= s[i] && s[i] <= 'Z') {
			newS = append(newS, s[i])
		}
	}
	// fmt.Println(string(newS))
	for i := 0; i < len(newS)/2; i++ {
		if newS[i] != newS[len(newS)-1-i] {
			return false
		}
	}

	return true
}
