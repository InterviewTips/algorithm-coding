package _348

import (
	"testing"
)

func TestZeroFilledSubarray(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int64
	}{
		{[]int{0, 0, 1, 0, 0, 0, 1}, 9},
		{[]int{0, 1, 0, 1, 0, 0, 1}, 5},
		{[]int{1, 1, 1, 1, 1}, 0},
		{[]int{0, 0, 0, 0, 0}, 15},
		{[]int{}, 0},
	}

	for _, test := range tests {
		got := zeroFilledSubarray(test.nums)
		if got != test.expect {
			t.Errorf("zeroFilledSubarray(%v) = %v, expect %v", test.nums, got, test.expect)
		}
	}
}
