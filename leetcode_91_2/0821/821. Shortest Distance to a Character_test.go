package _821

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_shortestToChar(t *testing.T) {
	assert.Equal(
		t,
		[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		shortestToChar("loveleetcode", 'e'),
	)
}
