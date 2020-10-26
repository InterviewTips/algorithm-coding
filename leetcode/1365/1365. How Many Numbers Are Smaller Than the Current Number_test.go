package _365

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	assert.Equal(
		t,
		[]int{4, 0, 1, 1, 3},
		smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}),
	)
}
