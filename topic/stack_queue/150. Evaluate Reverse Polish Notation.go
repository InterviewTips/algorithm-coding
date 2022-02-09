package stack_queue

import "strconv"

/*逆波兰求值
又是栈的应用，遇到运算符则运算

有效的算符包括 +、-、*、/
*/

func evalRPN(tokens []string) int {
	stack1 := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		// 运算符 取出两位进行运算 然后 push
		case "+":
			num1 := stack1[len(stack1)-1]
			num2 := stack1[len(stack1)-2]
			stack1 = stack1[:len(stack1)-2]
			num := num1 + num2
			stack1 = append(stack1, num)
		case "-":
			num1 := stack1[len(stack1)-1]
			num2 := stack1[len(stack1)-2]
			stack1 = stack1[:len(stack1)-2]
			num := num2 - num1 // 注意是后一个减前一个
			stack1 = append(stack1, num)
		case "*":
			num1 := stack1[len(stack1)-1]
			num2 := stack1[len(stack1)-2]
			stack1 = stack1[:len(stack1)-2]
			num := num1 * num2
			stack1 = append(stack1, num)
		case "/":
			num1 := stack1[len(stack1)-1]
			num2 := stack1[len(stack1)-2]
			stack1 = stack1[:len(stack1)-2]
			num := num2 / num1 // 注意是后一个除以前一个
			stack1 = append(stack1, num)
		default:
			// 转换为 int
			value, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			stack1 = append(stack1, value)
		}
	}

	return stack1[0]
}
