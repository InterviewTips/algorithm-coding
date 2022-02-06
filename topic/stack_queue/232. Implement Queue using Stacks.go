package stack_queue

// MyQueue 请你仅使用两个栈实现先入先出队列
type MyQueue struct {
	Stack1 []int // 用于 push
	Stack2 []int // 用于 pop
}

func Constructor() MyQueue {
	return MyQueue{
		Stack1: make([]int, 0),
		Stack2: make([]int, 0),
	}
}

func (m *MyQueue) Push(x int) {
	// push to stack1
	m.Stack1 = append(m.Stack1, x)
}

func (m *MyQueue) Pop() int {
	// 假定都是正常 pop 的
	if len(m.Stack2) == 0 {
		m.Stack2 = append(m.Stack2, m.Stack1...)
		// 清空 stack1
		m.Stack1 = make([]int, 0)
	}
	//res := m.Stack2[len(m.Stack2)-1]
	res := m.Stack2[0]
	// stack2 重新赋值
	m.Stack2 = m.Stack2[1:]
	return res
}

func (m *MyQueue) Peek() int {
	// 假定都是正常 peek 的
	if len(m.Stack2) == 0 {
		m.Stack2 = append(m.Stack2, m.Stack1...)
		// 清空 stack1
		m.Stack1 = make([]int, 0)
	}
	res := m.Stack2[0]
	return res
}

func (m *MyQueue) Empty() bool {
	if len(m.Stack1) == 0 && len(m.Stack2) == 0 {
		return true
	}

	return false
}
