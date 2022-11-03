package issue9

import (
	"log"
)

// 使用数组实现即可

type CustomStack struct {
	nums []int
	size int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		nums: make([]int, 0),
		size: maxSize,
	}
}

func (c *CustomStack) Push(x int) {
	if len(c.nums) >= c.size {
		log.Println("size overflow")
		return
	}

	c.nums = append(c.nums, x)
	return
}

func (c *CustomStack) Pop() int {
	if len(c.nums) <= 0 {
		return -1
	}
	length := len(c.nums)
	value := c.nums[length-1]
	c.nums = c.nums[:length-1]
	return value
}

func (c *CustomStack) Increment(k int, val int) {
	if k > len(c.nums) {
		k = len(c.nums)
	}
	for i := 0; i < k; i++ {
		c.nums[i] += val
	}
	return
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
