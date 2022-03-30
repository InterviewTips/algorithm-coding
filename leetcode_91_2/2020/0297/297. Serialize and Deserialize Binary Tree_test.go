package _297

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "nil,123,"
	res := strings.Split(s, ",")
	fmt.Printf("len(s) is %d, s is %v\n", len(res), res)
}
