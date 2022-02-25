package greedy

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// 身高高的排前面
	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		position := people[i][1] //插入的位置
		res = insertElement(res, position, people[i])
	}

	return res
}

func insertElement(data [][]int, position int, element []int) [][]int {
	left := data[:position]
	right := data[position:]
	var res [][]int
	res = append(res, left...)
	res = append(res, element)
	res = append(res, right...)

	return res
}
