package _347

import "container/heap"

type Num struct {
	Key   int
	Value int
}

type NumList []Num

func (n *NumList) Push(x interface{}) {
	*n = append(*n, x.(Num))
}

func (n *NumList) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	*n = old[:len(old)-1]
	return x
}

func (n NumList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NumList) Len() int           { return len(n) }
func (n NumList) Less(i, j int) bool { return n[i].Value < n[j].Value }

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &NumList{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, Num{
			Key:   key,
			Value: value,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).(Num).Key
	}
	return ret
}
