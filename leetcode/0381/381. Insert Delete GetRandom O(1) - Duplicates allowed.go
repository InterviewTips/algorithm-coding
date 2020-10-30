package _381

import (
	"math/rand"
	"time"
)

type RandomizedCollection struct {
	idx  map[int]map[int]string // 第一个 int 表示 val，第二个 int 表示 val 所处的索引，string 没有具体含义
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: map[int]map[int]string{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (r *RandomizedCollection) Insert(val int) bool {
	ids, has := r.idx[val]
	if !has {
		ids = map[int]string{}
		r.idx[val] = ids
	}
	ids[len(r.nums)] = "" // 记录位置索引
	r.nums = append(r.nums, val)
	return !has
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (r *RandomizedCollection) Remove(val int) bool {
	ids, has := r.idx[val]
	if !has {
		return false
	}
	var i int
	for id := range ids {
		i = id // 第一个等于 val 的索引位置
		break
	}
	n := len(r.nums)
	r.nums[i] = r.nums[n-1]       // 列表最后一个元素赋值给待删除的元素
	delete(ids, i)                // 删除 ids 中 key 为 i 的 key
	delete(r.idx[r.nums[i]], n-1) // 删除原先最后一个元素的索引
	if i < n-1 {                  // 删除的不是最后一个元素，所以需要重建索引
		r.idx[r.nums[i]][i] = ""
	}
	if len(ids) == 0 { // 如果 ids 删除之后没有 key 了，就删除 r.idx 中 key 为这个 val 的 key
		delete(r.idx, val)
	}
	r.nums = r.nums[:n-1] // 删除列表最后一个元素
	return true
}

/** Get a random element from the collection. */
func (r *RandomizedCollection) GetRandom() int {
	rand.Seed(time.Now().UnixNano()) // 注：这里不 seed 其实达不到随机，用了这个函数之后 运行时间会变成 180ms 左右，没用 40ms 左右
	return r.nums[rand.Intn(len(r.nums))]
}
