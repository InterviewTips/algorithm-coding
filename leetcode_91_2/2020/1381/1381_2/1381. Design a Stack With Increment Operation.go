package _381

type CustomStack struct {
	s      []int
	a      []int
	offset int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		s:      make([]int, maxSize),
		a:      make([]int, maxSize),
		offset: 0,
	}
}

func (c *CustomStack) Push(x int) {
	if c.offset >= len(c.s) {
		return
	}
	c.s[c.offset] = x
	c.offset++
	return
}

func (c *CustomStack) Pop() int {
	if c.offset > 0 {
		add := c.a[c.offset-1]
		res := c.s[c.offset-1] + add
		if c.offset-1 > 0 {
			c.a[c.offset-2] += add
		}
		c.a[c.offset-1] = 0
		c.offset--
		return res
	}
	return -1
}

func (c *CustomStack) Increment(k int, val int) {
	if c.offset < k {
		k = c.offset
	}
	if k > 0 {
		c.a[k-1] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
