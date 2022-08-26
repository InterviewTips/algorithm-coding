package main

import "fmt"

func main() {
	var n int
	var sum int
	var oneNum int
	for {
		count, _ := fmt.Scan(&n)
		if count == 0 {
			break // return 也可以
		}
		sum = 0
		for i := 0; i < n; i++ {
			fmt.Scan(&oneNum)
			sum += oneNum
		}
		fmt.Println(sum)
	}
}
