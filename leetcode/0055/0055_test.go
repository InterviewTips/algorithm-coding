package _055

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCanJump(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}), "must be true")
	assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}), "must be false")
}
