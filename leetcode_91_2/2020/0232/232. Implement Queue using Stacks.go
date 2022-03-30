package _232

type MyQueue struct {
	stack1 []int
	stack2 []int
}

// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) { // O(1)
	m.stack1 = append(m.stack1, x)
	return
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int { //
	if len(m.stack2) != 0 {
	} else {
		for len(m.stack1) != 0 {
			m.stack2 = append(m.stack2, m.stack1[len(m.stack1)-1])
			m.stack1 = m.stack1[:len(m.stack1)-1]
		}
	}
	res := m.stack2[len(m.stack2)-1]
	m.stack2 = m.stack2[:len(m.stack2)-1]
	return res
}

/** Get the front element. */
func (m *MyQueue) Peek() int { //
	if len(m.stack2) != 0 {
	} else {
		for len(m.stack1) != 0 {
			m.stack2 = append(m.stack2, m.stack1[len(m.stack1)-1])
			m.stack1 = m.stack1[:len(m.stack1)-1]
		}
	}
	return m.stack2[len(m.stack2)-1]
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool { // O(1)
	return len(m.stack1) == 0 && len(m.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
