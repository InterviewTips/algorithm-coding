package _155

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		obj := Constructor()
		obj.Push(-2)
		obj.Push(0)
		obj.Push(-3)
		t.Logf("getMin is %d\n", obj.GetMin())
		obj.Pop()
		t.Logf("top is %d\n", obj.Top())
		t.Logf("getMin is %d\n", obj.GetMin())
	})
}
