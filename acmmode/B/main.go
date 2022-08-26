package main

import "fmt"

func main() {
	var t int
	var a int
	var b int
	n, _ := fmt.Scan(&t)
	if n == 0 {
		panic("输入 error")
	}
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
