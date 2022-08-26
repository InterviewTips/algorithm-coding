package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var count int
	fmt.Scan(&count)
	lines := make([]string, 0)
	for i := 0; i < count; i++ {
		var word string
		fmt.Scan(&word)
		lines = append(lines, word)
	}
	sort.Strings(lines)
	//sort.SliceStable(lines, func(i, j int) bool {
	//	return lines[i] < lines[j]
	//})
	fmt.Println(strings.Join(lines, " "))
}
