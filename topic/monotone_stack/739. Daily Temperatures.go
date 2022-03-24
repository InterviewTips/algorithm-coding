package monotone_stack

import "log"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	stack := make([]int, 0) // 栈顶到栈底 单调递增

	stack = append(stack, 0) // 1 <= temperatures.length <= 105
	for i := 1; i < len(temperatures); i++ {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 计算 res
			res[stack[len(stack)-1]] = i - stack[len(stack)-1] // 第 i 天后
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

		//if len(stack) == 0 {
		//	stack = append(stack, i)
		//	continue
		//}
		//if temperatures[i] <= temperatures[stack[len(stack)-1]] {
		//	stack = append(stack, i)
		//} else { // > 不断弹出
		//}
	}

	log.Println(res)

	return res
}
