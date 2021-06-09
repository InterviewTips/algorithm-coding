package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	c := Constructor()
	c.AppendTail(1)
	c.AppendTail(2)
	assert.Equal(t, stack{1, 2}, c.appendStack)
}

func TestCQueue_DeleteHead(t *testing.T) {
	c := Constructor()
	c.AppendTail(1)
	c.AppendTail(2)
	assert.Equal(t, 1, c.DeleteHead())
	assert.Equal(t, 2, c.DeleteHead())
}

func TestMix(t *testing.T) {
	c := Constructor()
	c.AppendTail(3)
	t.Log(c.DeleteHead())
	t.Log(c.DeleteHead())
	t.Log(c.DeleteHead())
}
