package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		words := strings.Split(inputs.Text(), ",")
		sort.Strings(words)
		fmt.Println(strings.Join(words, ","))
	}
}
