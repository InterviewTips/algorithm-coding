package _022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"), "must be true")
	// t.Logf("isValid is %v\n", isValid("()"))
}
