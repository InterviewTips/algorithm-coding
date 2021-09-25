package array

import (
	//"log"
	"math"
)

// 在一排树中，第 i 棵树产生 tree[i] 型的水果。
//你可以从你选择的任何树开始，然后重复执行以下步骤：
//
//把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
//移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
//请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。
//
//你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
//
//用这个程序你能收集的水果树的最大总量是多少？
//
//示例 1：
//
//输入：[1,2,1]
//输出：3
//解释：我们可以收集 [1,2,1]。
//示例 2：
//
//输入：[0,1,2,2]
//输出：3
//解释：我们可以收集 [1,2,2]
//如果我们从第一棵树开始，我们将只能收集到 [0, 1]。
//示例 3：
//
//输入：[1,2,3,2,2]
//输出：4
//解释：我们可以收集 [2,3,2,2]
//如果我们从第一棵树开始，我们将只能收集到 [1, 2]。
//示例 4：
//
//输入：[3,3,3,1,2,1,1,2,3,3,4]
//输出：5
//解释：我们可以收集 [1,2,1,1,2]
//如果我们从第一棵树或第八棵树开始，我们将只能收集到 4 棵水果树。
//
//提示：
//
//1 <= tree.length <= 40000
//0 <= tree[i] < tree.length
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fruit-into-baskets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func totalFruit(fruits []int) int {

	result := math.MinInt32
	left := 0
	// 维护一个 map 进行判断
	fruitsMap := make(map[int]int) // 计算出现次数
	for right := 0; right < len(fruits); right++ {
		singleFruit := fruits[right]
		fruitsMap[singleFruit] = fruitsMap[singleFruit] + 1
		for len(fruitsMap) >= 3 { // 移除元素
			fruitsMap[fruits[left]] = fruitsMap[fruits[left]] - 1
			if fruitsMap[fruits[left]] == 0 {
				// log.Println("删除元素", fruits[left])
				delete(fruitsMap, fruits[left]) // 这里删除之后，上面的 for 条件自然不成立
				// left++
				// break
			}
			left++
		}

		count := right - left + 1
		if count > result {
			result = count
		}
	}

	return result
}

func totalFruit1(fruits []int) int {
	result := math.MinInt32
	left := 0
	// 维护一个 map 进行判断
	fruitsMap := make(map[int]int) // 计算出现次数
	for right := 0; right < len(fruits); right++ {
		singleFruit := fruits[right]
		_, ok := fruitsMap[singleFruit]
		if ok {
			// log.Println("存在，", singleFruit)
			fruitsMap[singleFruit] = fruitsMap[singleFruit] + 1
			// 计数
			count := right - left + 1
			if count > result {
				result = count
			}
		} else {
			if len(fruitsMap) >= 2 {
				// log.Println(">=2", singleFruit)
				// 移动元素 一个元素
				for fruitsMap[fruits[left]] != 0 {
					// log.Println(
					//	"减掉",
					//	fruits[left],
					//	"当前数量",
					//	fruitsMap[fruits[left]],
					//	"删除后数量",
					//	fruitsMap[fruits[left]]-1,
					//)
					fruitsMap[fruits[left]] = fruitsMap[fruits[left]] - 1
					if fruitsMap[fruits[left]] == 0 {
						// log.Println("删除元素", fruits[left])
						delete(fruitsMap, fruits[left])
						left++
						break // 已经移除了一个元素 可以退出
					}
					left++
				}
				// 新进元素
				fruitsMap[singleFruit] = fruitsMap[singleFruit] + 1
				// log.Println(">=2 新进元素", singleFruit)
			} else { // 计数
				// log.Println("新进元素", singleFruit)
				fruitsMap[singleFruit] = fruitsMap[singleFruit] + 1
				count := right - left + 1
				if count > result {
					result = count
				}
			}
		}
	}

	return result
}
