package _394

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_decodeString(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
	assert.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
	assert.Equal(t, "abccdcdcdxyz", decodeString("abc3[cd]xyz"))
}
