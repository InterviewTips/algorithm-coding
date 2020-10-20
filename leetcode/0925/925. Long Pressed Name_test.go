package _925

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_isLongPressedName(t *testing.T) {
	assert.Equal(
		t,
		true,
		isLongPressedName("alex", "aaleex"),
	)
	assert.Equal(
		t,
		false,
		isLongPressedName("saeed", "ssaaedd"),
	)
	assert.Equal(
		t,
		true,
		isLongPressedName("leelee", "lleeelee"),
	)
	assert.Equal(
		t,
		true,
		isLongPressedName("laiden", "laiden"),
	)
	assert.Equal(
		t,
		false,
		isLongPressedName("alex", "aaleelx"),
	)
	assert.Equal(
		t,
		false,
		isLongPressedName("dfuyalc", "fuuyallc"),
	)
	assert.Equal(
		t,
		true,
		isLongPressedName("vtkgn", "vttkgnn"),
	)
}
