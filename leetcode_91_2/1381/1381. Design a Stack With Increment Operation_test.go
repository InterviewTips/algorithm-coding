package _381

import (
	"testing"
)

func TestCustomStack(t *testing.T) {
	c := Constructor(3)
	c.Push(1)
	c.Push(2)
	c.Pop()
	c.Push(2)
	c.Push(3)
	c.Push(4)
	c.Increment(5, 100)
	c.Increment(2, 100)
	c.Pop()
	c.Pop()
	c.Pop()
	c.Pop()
}
