package main

import "fmt"

func main() {
	var a int
	var b int
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 { // 退出条件
			break
		}
		fmt.Printf("%d\n", a+b) // 输出换行
	}
}
