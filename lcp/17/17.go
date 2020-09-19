package main

func calculate(s string) int {
	//"A" 运算：使 x = 2 * x + y
	//"B" 运算：使 y = 2 * y + x
	x := 1
	y := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			x = 2*x + y
		case 'B':
			y = 2*y + x
		}
	}
	return x + y
}
