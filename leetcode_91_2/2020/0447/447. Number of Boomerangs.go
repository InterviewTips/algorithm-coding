package _447

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		distance := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i != j {
				d := distanceCalc(points[i], points[j])
				v, ok := distance[d]
				fmt.Printf("d is %v\n", d)
				if ok { // 存在
					res += 2 * v // (v+1) * v - [v * (v-1)] = 2 * v
					distance[d] += 1
				} else { // 不存在则写入
					distance[d] = 1
				}
			}
		}
	}
	return res
}

func distanceCalc(a, b []int) int {
	return ((a[0] - b[0]) * (a[0] - b[0])) + ((a[1] - b[1]) * (a[1] - b[1]))
}

func numberOfBoomerangs1(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		distance := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i != j {
				d := distanceCalc(points[i], points[j])
				distance[d] += 1
			}
		}
		for _, v := range distance {
			res += v * (v - 1)
		}
	}
	return res
}
