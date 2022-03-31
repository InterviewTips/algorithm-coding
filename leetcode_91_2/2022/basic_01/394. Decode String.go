package basic_01

import (
	"log"
	"strconv"
)

func decodeString(s string) string {
	// 2[abc]
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			stack = append(stack, s[i])
		case s[i] == '[':
			stack = append(stack, s[i])
		case s[i] == ']':
			// 取出前面的字母串 [
			index := len(stack) - 1
			for index >= 0 && stack[index] != '[' {
				index--
			}

			//stack[index] == '['
			// 要保证 letter 不会被修改到 所以使用 copy
			letter := make([]byte, len(stack[index+1:]))
			copy(letter, stack[index+1:])
			//letter := stack[index+1:] 错误写法
			// 例如 3[acc
			// 此时 letter = acc, num = 3
			// 经过第一次修改 stack 变为 acc
			// 此时 stack[2] = c, 所以 letter 的第一个元素会变为 c, letter = ccc
			// 最后产生错误结果 acccccccc
			log.Println("letter is", string(letter))

			// 取出前面的数字 !letter
			numIndex := index - 1 // 从倒数第一个数字开始 index - 1
			for numIndex >= 0 && (stack[numIndex] >= '0' && stack[numIndex] <= '9') {
				numIndex--
			}

			numStr := string(stack[numIndex+1 : index])
			//log.Println("num is", numStr)
			num, _ := strconv.Atoi(numStr)

			// 去掉 stack 对应的值
			stack = stack[:numIndex+1]

			// 加入生成的字符串
			for i := 0; i < num; i++ {
				//log.Println("for letter is ", string(letter))
				stack = append(stack, letter...)
			}
		default:
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
