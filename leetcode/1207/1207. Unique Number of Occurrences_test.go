package _207

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_uniqueOccurrences(t *testing.T) {
	assert.Equal(t, true, uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
