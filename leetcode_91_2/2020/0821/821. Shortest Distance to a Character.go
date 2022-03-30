package _821

func shortestToChar(S string, C byte) []int {
	a := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			a = append(a, i)
		}
	}
	res := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			//res = append(res, 0)
			res[i] = 0
		} else {
			min := len(S)
			for j := 0; j < len(a); j++ {
				if abs(a[j], i) < min {
					min = abs(a[j], i)
				}
			}
			//res = append(res, min)
			res[i] = min
		}
	}
	return res
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}
