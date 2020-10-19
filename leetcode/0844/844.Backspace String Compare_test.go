package _844

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_backspaceCompare(t *testing.T) {
	assert.Equal(
		t,
		false,
		backspaceCompare("122", "123"),
	)
	assert.Equal(
		t,
		true,
		backspaceCompare("123", "123"),
	)
	assert.Equal(
		t,
		true,
		backspaceCompare("1#3", "3"),
	)
	assert.Equal(
		t,
		true,
		backspaceCompare("1##", "#"),
	)
}
