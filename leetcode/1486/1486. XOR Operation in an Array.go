package _486

func xorOperation(n int, start int) int {
	// start + 2 * i
	if n == 0 {
		return 0
	}
	if n == 1 {
		return start
	}
	res := start
	for i := 1; i < n; i++ {
		x := start + 2*i
		res = res ^ x
	}
	return res
}
