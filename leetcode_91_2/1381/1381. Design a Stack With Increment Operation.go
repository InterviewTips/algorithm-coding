package _381

type CustomStack struct {
	s      []int
	offset int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{s: make([]int, maxSize), offset: 0}
}

func (c *CustomStack) Push(x int) {
	//fmt.Printf("len(c.s) is %v\n", len(c.s))
	if c.offset >= len(c.s) {
		return
	}
	c.s[c.offset] = x
	c.offset++
	return
}

func (c *CustomStack) Pop() int {
	if c.offset > 0 {
		res := c.s[c.offset-1]
		c.offset--
		return res
	}
	return -1
}

func (c *CustomStack) Increment(k int, val int) {
	if c.offset < k {
		k = c.offset
	}
	for i := 0; i < k; i++ {
		c.s[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
