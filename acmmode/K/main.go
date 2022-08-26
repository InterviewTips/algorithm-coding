package main

import "fmt"

func main() {
	var a int64
	var b int64
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}
		fmt.Println(a + b)
	}

}
