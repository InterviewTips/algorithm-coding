package string

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	a := strings.Split(s, " ")
	b := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == "" {
			continue
		}
		b = append(b, a[i])
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-(i+1)] = b[len(b)-(i+1)], b[i]
	}
	fmt.Println(b)
	c := strings.Join(b, " ")
	// fmt.Println(b)
	return c
}
