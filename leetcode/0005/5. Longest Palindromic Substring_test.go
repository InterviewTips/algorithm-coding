package _005

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
}
