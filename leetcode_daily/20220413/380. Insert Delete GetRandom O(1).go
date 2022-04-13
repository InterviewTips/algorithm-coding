package _0220413

import "math/rand"

type RandomizedSet struct {
	data map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{data: map[int]bool{}}
}

func (r *RandomizedSet) Insert(val int) bool {
	value := r.data[val]
	r.data[val] = true
	return !value
}

func (r *RandomizedSet) Remove(val int) bool {
	value := r.data[val]
	// todo: 变长数组 用最后一个元素替换即可
	delete(r.data, val)
	return value
}

func (r *RandomizedSet) GetRandom() int {
	nums := make([]int, 0)
	for k, _ := range r.data {
		nums = append(nums, k)
	}

	return nums[rand.Intn(len(nums))]
}
