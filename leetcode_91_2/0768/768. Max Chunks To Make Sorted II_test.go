package _768

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_slice(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 3
	t.Logf("a is %v, b is %v\n", a, b)
}

func Test_maxChunksToSorted(t *testing.T) {
	assert.Equal(
		t,
		1,
		maxChunksToSorted([]int{5, 4, 3, 2, 1}),
	)
	assert.Equal(
		t,
		4,
		maxChunksToSorted([]int{2, 1, 3, 4, 4}),
	)
}
