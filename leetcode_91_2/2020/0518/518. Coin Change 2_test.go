package _518

import "testing"

func TestChange(t *testing.T) {
	t.Logf("change is %v\n", change(5, []int{1, 2, 5}))
	t.Logf("change is %v\n", change(3, []int{2}))
	t.Logf("change is %v\n", change(10, []int{10}))

	t.Logf("change1 is %v\n", change1(5, []int{1, 2, 5}))
	t.Logf("change1 is %v\n", change1(3, []int{2}))
	t.Logf("change1 is %v\n", change1(10, []int{10}))
}
