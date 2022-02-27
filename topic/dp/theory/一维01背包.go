package theory

import "log"

func oneDimensional01Bag(weight, value []int, bagWeight int) int {
	// 表示背包容量为 j 时，可以获取的最大价值总和
	// res[j] = getMax(res[j], res[j-weight[i]]+value[i])
	// res[j]
	// res[j-weight[i]] + value[i]
	res := make([]int, bagWeight+1)
	res[0] = 0 // 背包容量为 0 所以为 0

	//for i := 0; i < len(weight); i++ { // 错误写法: j 从前到后遍历会多次取相同的物品
	//	for j := weight[i]; j <= bagWeight; j++ { // 保证 j - weight[i] >= 0
	//		log.Println("weight[i]", weight[i], "i:", i,"j:",j, "res is:", res)
	//		res[j] = getMax(res[j], res[j-weight[i]]+value[i])
	//	}
	//}
	//log.Println("last res:",res)

	// 错误写法：如果先遍历背包容量，则每个背包最后其实只会放入一个物品而已
	// res is [0 15 15 20 30]
	//for j := bagWeight; j >= 1; j-- { // 保证 j - weight[i] >= 0
	//	for i := 0; i < len(weight); i++ {
	//		if j - weight[i] < 0 {
	//			continue
	//		}
	//		log.Println("i:", i, "前 weight[i]", weight[i], "背包容量:", j, "j-weight[i]=", j-weight[i], "res is:", res)
	//		res[j] = getMax(res[j], res[j-weight[i]]+value[i])
	//		log.Println("i:", i, "后 weight[i]", weight[i], "背包容量:", j, "j-weight[i]=", j-weight[i], "res is:", res)
	//	}
	//}
	//log.Println("last res:", res)

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- { // 保证 j - weight[i] >= 0
			//if j - weight[i] < 0 { 通过上面条件判断即可
			//	continue
			//}
			log.Println("i:", i, "前 weight[i]", weight[i], "背包容量:", j, "j-weight[i]=", j-weight[i], "res is:", res)
			res[j] = getMax(res[j], res[j-weight[i]]+value[i])
			log.Println("i:", i, "后 weight[i]", weight[i], "背包容量:", j, "j-weight[i]=", j-weight[i], "res is:", res)
		}
	}
	log.Println("last res:", res)

	return res[bagWeight]
}
