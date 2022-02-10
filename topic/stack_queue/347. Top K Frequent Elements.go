package stack_queue

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range count {
		heap.Push(h, KV{key: key, value: value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		key := heap.Pop(h).(KV).key // k 个中最小的
		res[k-i-1] = key            // 所以要 k-i-1 倒序放入 因为 res 是高频在前
	}
	return res
}

type KV struct {
	key   int
	value int
}

// IHeap 构建小顶堆
type IHeap []KV

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool { // 小顶堆 堆头为最小元素
	return (*h)[i].value < (*h)[j].value
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
