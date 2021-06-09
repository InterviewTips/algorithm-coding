package main

type stack []int

type CQueue struct {
	appendStack stack
	popStask    stack
}

func Constructor() CQueue {
	return CQueue{
		appendStack: make(stack, 0),
		popStask:    make(stack, 0),
	}
}

// 添加到尾部
func (c *CQueue) AppendTail(value int) {
	c.appendStack = append(c.appendStack, value)
}

// 取出一个元素
func (c *CQueue) DeleteHead() int {
	var res int
	// 先看看 popStack 有没有数据，有则直接返回
	if len(c.popStask) != 0 {
		res = c.popStask[len(c.popStask)-1]
		// 删除这一个元素
		c.popStask = c.popStask[:len(c.popStask)-1]
		return res
	}
	// popStack 没有元素 则看看 appendStack 中的 没有则直接返回 panic 吧
	if len(c.appendStack) == 0 {
		return -1
	}
	// 先将 appendStack 一个一个入到 popStack 中
	for i := len(c.appendStack) - 1; i >= 0; i-- {
		c.popStask = append(c.popStask, c.appendStack[i])
	}
	// 清空 appendStack
	c.appendStack = c.appendStack[:0]

	res = c.popStask[len(c.popStask)-1]
	// 删除这一个元素
	c.popStask = c.popStask[:len(c.popStask)-1]
	return res
}
