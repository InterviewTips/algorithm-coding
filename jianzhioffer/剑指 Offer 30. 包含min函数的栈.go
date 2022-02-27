package jianzhioffer

import "math"

type MinStack struct {
	min   [2]int // 需要记录下第一个出现最小值的索引
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   [2]int{0, math.MaxInt64},
		stack: make([]int, 0),
	}
}

func (m *MinStack) Push(x int) {
	if x < m.min[1] {
		m.min = [2]int{len(m.stack), x}
	}

	m.stack = append(m.stack, x)
}

// todo: 这里 pop 的操作是 O(n) 如果要改为 O(1) 需要再多使用一个栈 保证栈中的栈顶元素是最小的，这里还有一个点需要注意 如果是同样是最小值
// 也需要 push 进入辅助栈
func (m *MinStack) Pop() {
	x := m.stack[len(m.stack)-1]
	if x == m.min[1] && len(m.stack)-1 == m.min[0] {
		// 更新最小值
		minValue, index := m.getMinIndexValue(m.stack[:len(m.stack)-1])
		m.min = [2]int{minValue, index}
	}
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) getMinIndexValue(nums []int) (int, int) {
	minValue := math.MaxInt64
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < minValue {
			minValue = nums[i]
			index = i
		}
	}

	return index, minValue
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) Min() int {
	return m.min[1]
}
