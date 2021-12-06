package string

import "fmt"

func reverseString(s []byte) {
	fmt.Println(s)
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-(i+1)]
		s[len(s)-(i+1)] = tmp
	}
	fmt.Println(s)
}
