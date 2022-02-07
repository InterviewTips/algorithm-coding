package stack_queue

type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func stackConstructor() MyStack {
	return MyStack{
		Queue1: make([]int, 0),
		Queue2: make([]int, 0),
	}
}

func (m *MyStack) Push(x int) {
	if len(m.Queue2) != 0 {
		m.Queue2 = append(m.Queue2, x)
		return
	}
	m.Queue1 = append(m.Queue1, x)
}

func (m *MyStack) Pop() int {
	if len(m.Queue2) == 0 {
		m.Queue2 = append(m.Queue2, m.Queue1...)
		// 清空 queue1
		m.Queue1 = make([]int, 0)
	}
	res := m.Queue2[len(m.Queue2)-1]
	m.Queue2 = m.Queue2[:len(m.Queue2)-1]
	return res
}

func (m *MyStack) Top() int {
	if len(m.Queue2) == 0 {
		m.Queue2 = append(m.Queue2, m.Queue1...)
		// 清空 queue1
		m.Queue1 = make([]int, 0)
	}
	res := m.Queue2[len(m.Queue2)-1]
	return res
}

func (m *MyStack) Empty() bool {
	if len(m.Queue1) == 0 && len(m.Queue2) == 0 {
		return true
	}

	return false
}
