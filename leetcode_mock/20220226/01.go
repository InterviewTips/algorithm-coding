package _0220226

// https://leetcode-cn.com/problems/airplane-seat-assignment-probability/
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}

	return 0.5
}
