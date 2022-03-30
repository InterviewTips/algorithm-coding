package _232

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_MyQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, true, q.Empty())
}
