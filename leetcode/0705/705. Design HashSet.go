package _705

type MyHashSet struct {
	nums []bool
}

// 注: 0 <= key <= 10e6
func Constructor() MyHashSet {
	return MyHashSet{
		nums: make([]bool, 10e6+1),
	}
}

func (m *MyHashSet) Add(key int) {
	m.nums[key] = true
}

func (m *MyHashSet) Remove(key int) {
	m.nums[key] = false
}

func (m *MyHashSet) Contains(key int) bool {
	return m.nums[key]
}
