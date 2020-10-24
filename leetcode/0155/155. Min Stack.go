package _155

type MinStack struct {
	min   int
	nodes []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   1 << 31,
		nodes: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	// 维护 min
	if x < this.min {
		this.min = x
	}
	this.nodes = append(this.nodes, x)
}

func (this *MinStack) Pop() {
	if len(this.nodes) == 0 {
		return
	}
	// 维护 min
	if this.nodes[len(this.nodes)-1] == this.min {
		min := 1 << 31
		for i := 0; i < len(this.nodes)-1; i++ {
			if this.nodes[i] < min {
				min = this.nodes[i]
			}
		}
		this.min = min
	}
	this.nodes = this.nodes[:len(this.nodes)-1]
}

func (this *MinStack) Top() int {
	if len(this.nodes) == 0 {
		return 1 << 31
	}
	return this.nodes[len(this.nodes)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
