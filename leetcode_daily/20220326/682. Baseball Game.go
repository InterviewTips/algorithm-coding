package _0220326

import "strconv"

func calPoints(ops []string) int {
	// 模拟堆栈操作
	stack := make([]int, 0)
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			// 删除上一个
			stack = stack[:len(stack)-1]
		case "D":
			// 上一个 double
			stack = append(stack, 2*stack[len(stack)-1])
		case "+":
			// 上两个相加
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			value, _ := strconv.Atoi(ops[i])
			stack = append(stack, value)
		}
	}

	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}

	return sum
}
