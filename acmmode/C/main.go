package main

import "fmt"

func main() {
	var a int
	var b int
	for {
		num, _ := fmt.Scan(&a, &b)
		if num == 0 || (a == 0 && b == 0) {
			break
		}
		fmt.Println(a + b)
	}

}
