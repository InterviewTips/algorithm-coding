package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var n int
	var sum int
	var oneNum int
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		sum = 0
		for i := 0; i < n; i++ {
			fmt.Scan(&oneNum)
			sum += oneNum
		}
		fmt.Println(sum)
	}
}
