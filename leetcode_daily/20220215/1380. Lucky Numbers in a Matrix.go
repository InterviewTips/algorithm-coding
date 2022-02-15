package _0220215

import (
	"math"
)

// 行最小 列最大
func luckyNumbers(matrix [][]int) []int {
	res := make([]int, 0)
	x := len(matrix) // 行
	//y := len(matrix[0]) // 列

	// 取出每一行的最小值 索引
	xMinIndexs := make([][2]int, x)
	for i := 0; i < x; i++ {
		xMinIndexs[i] = [2]int{i, getMin(matrix[i])}
	}

	// 取出每一行最小值的对应那一列的最大值索引
	yMaxIndexs := make([][2]int, x)
	yMaxIndexsTable := make(map[int][2]int)
	for i := 0; i < x; i++ {
		y := xMinIndexs[i][1]
		// 取一下 map
		value, ok := yMaxIndexsTable[y]
		if ok {
			yMaxIndexs[i] = value
		} else {
			yData := make([]int, 0)
			for j := 0; j < x; j++ {
				yData = append(yData, matrix[j][y])
			}
			yMaxIndexs[i] = [2]int{getMax(yData), y}
			// set to map
			yMaxIndexsTable[y] = yMaxIndexs[i]
		}
		if yMaxIndexs[i] == xMinIndexs[i] {
			res = append(res, matrix[xMinIndexs[i][0]][xMinIndexs[i][1]])
		}
	}
	//log.Println("xIndex", xMinIndexs)
	//log.Println("yIndex", yMaxIndexs)

	return res
}

func getMin(data []int) int {
	min := math.MaxInt64
	var minIndex int
	for i := 0; i < len(data); i++ {
		if data[i] < min {
			min = data[i]
			minIndex = i
		}
	}

	return minIndex
}

func getMax(data []int) int {
	max := math.MinInt64
	var maxIndex int
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
			maxIndex = i
		}
	}

	return maxIndex
}
