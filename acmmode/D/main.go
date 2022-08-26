package main

import "fmt"

func main() {
	var n int
	var sum int
	var oneNum int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum = 0
		for i := 0; i < n; i++ {
			fmt.Scan(&oneNum)
			sum += oneNum
		}
		fmt.Println(sum)
	}
}
