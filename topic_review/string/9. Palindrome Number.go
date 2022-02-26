package string

import "fmt"

func isPalindrome(x int) bool {
	value := fmt.Sprintf("%v", x)
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-1-i] {
			return false
		}
	}

	return true
}
