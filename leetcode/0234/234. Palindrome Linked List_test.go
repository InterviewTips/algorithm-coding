package _234

import (
	"testing"

	"github.com/bmizerany/assert"

	"algorithm/template"
)

func Test_isPalindrome(t *testing.T) {
	assert.Equal(
		t,
		false,
		isPalindrome(template.NewLinkedList([]int{1, 2})),
	)
	assert.Equal(
		t,
		true,
		isPalindrome(template.NewLinkedList([]int{1, 2, 2, 1})),
	)
}
