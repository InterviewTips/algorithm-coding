package _069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 2, mySqrt(4))
}
