package _052

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaxSatisfied(t *testing.T) {
	assert.Equal(
		t,
		16,
		maxSatisfied(
			[]int{1, 0, 1, 2, 1, 1, 7, 5},
			[]int{0, 1, 0, 1, 0, 1, 0, 1},
			3),
	)
}
