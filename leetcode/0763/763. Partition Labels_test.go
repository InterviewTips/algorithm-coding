package _763

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_partitionLabels(t *testing.T) {
	assert.Equal(
		t,
		[]int{9, 7, 8},
		partitionLabels("ababcbacadefegdehijhklij"),
	)
}
