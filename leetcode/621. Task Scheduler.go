package leetcode

import (
	"sort"
)

type task struct {
	name byte
	num  int
	n    int // 采用了就要开始计算 没采用就不用
}

func leastInterval(tasks []byte, n int) int {
	// 只需要计算最短时间即可
	data := make(map[byte]*task)
	for i := 0; i < len(tasks); i++ {
		item := tasks[i]
		v, ok := data[item]
		if !ok {
			data[item] = &task{
				name: item,
				num:  1,
				n:    0,
			}
			continue
		}
		v.num++
	}

	res := make([]*task, 0)
	for _, v := range data {
		res = append(res, v)
	}

	// sort
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].num > res[j].num
	})

	// 先消费 num 多的
	// 开始计算最短时间
	count := 0
	for {
		if len(data) == 0 {
			break
		}
		var choice byte
		for _, v := range res {
			//log.Printf("v is %+v\n", v)
			if v.n == 0 && v.num != 0 { // 采用
				//log.Printf("采用 %+v\n", v)
				choice = v.name
				v.num--
				v.n = n
				if v.num == 0 { // delete
					delete(data, v.name)
				}
				// 可能出现 有相同数量的存在 (即会出现 nums[0]没有永远大于num[1...] 的情况) 所以进行重新排序
				sort.SliceStable(res, func(i, j int) bool {
					return res[i].num > res[j].num
				})
				goto nSub
			}
		}
	nSub:
		// 所有非 0 的且没有被采用的 n--
		for _, v := range res {
			//log.Printf("range %+v\n", v)
			if v.name != choice && v.n > 0 {
				//log.Printf("n-- %+v\n",  v)
				v.n--
			}
		}
		count++ // 计算时间
	}

	return count
}
