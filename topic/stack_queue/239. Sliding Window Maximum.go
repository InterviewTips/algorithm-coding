package stack_queue

func maxSlidingWindow(nums []int, k int) []int {
	// 滑动窗口
	res := make([]int, 0)
	queue := monotonousQueue{nums: make([]int, 0)}
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}
	res = append(res, queue.front())
	for i := k; i < len(nums); i++ {
		queue.pop(nums[i-k])
		queue.push(nums[i])
		res = append(res, queue.front())
	}

	return res
}

//monotonousQueue 单调队列 队头最大
type monotonousQueue struct {
	nums []int
}

func (m *monotonousQueue) push(value int) {
	for !m.empty() && value > m.back() {
		m.popBack()
	}
	m.nums = append(m.nums, value)
}

func (m *monotonousQueue) pop(value int) {
	if !m.empty() && value == m.front() {
		m.popFront()
	}
}

func (m *monotonousQueue) empty() bool {
	return len(m.nums) == 0
}

func (m *monotonousQueue) popFront() {
	m.nums = m.nums[1:]
}

func (m *monotonousQueue) popBack() {
	m.nums = m.nums[:len(m.nums)-1]
}

func (m *monotonousQueue) front() int {
	return m.nums[0]
}

func (m *monotonousQueue) back() int {
	return m.nums[len(m.nums)-1]
}
