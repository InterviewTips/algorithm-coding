package main

import "fmt"

var stack1 []int // push
var stack2 []int // pop

func Push(node int) {
	stack1 = append(stack1, node)
	return
}

func Pop() int {
	if len(stack2) == 0 && len(stack1) == 0 {
		panic("can not pop")
	}
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	res := stack2[len(stack2)-1]
	fmt.Printf("res is %v\n", res)
	stack2 = stack2[:len(stack2)-1]
	fmt.Printf("stack2 is %v\n", stack2)
	return res
}
