package _022

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else { // ) ] }
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (s[i] == ')' && top == '(') ||
				(s[i] == ']' && top == '[') ||
				(s[i] == '}' && top == '{') {
				// pass
			} else {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
