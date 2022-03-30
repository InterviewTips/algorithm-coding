package _438

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	r := findAnagrams("cbaebabacd", "abc")
	fmt.Printf("r is %v\n", r)
}
