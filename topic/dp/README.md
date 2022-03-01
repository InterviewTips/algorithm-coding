<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Dynamic programming(dp)](#dynamic-programmingdp)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Dynamic programming(dp)

- 动态规划

```
递归问题 -> 重叠子问题&最优子结构 > 记忆化搜索(自顶向下解决问题)
                             > 动态规划(自下向上解决问题)
```

- 五步曲

```
1 确定dp数组（dp table）以及下标的含义
2 确定递推公式
3 dp数组如何初始化
4 确定遍历顺序
5 举例推导dp数组
```


- 遍历顺序的重要性

1、01 背包问题中，一维数组需要将物品放到外层，背包容量放里层，且容量需要从大到小，才不会进行重复计算

2、完全背包问题中，计算排列和组合的循环里外层是不同的

> 组合


```go
func change(amount int, coins []int) int {
	res := make([]int, amount+1)
	// res[j] amount = j 有多少种方式
	res[0] = 1
	for i := 0; i < len(coins); i++ { // 物品在外层
		for j := coins[i]; j <= amount; j++ { // 容量在里层
			res[j] += res[j-coins[i]]
		}
		//log.Println(res)
	}

	return res[amount]
}
```

> 排列

```go
func combinationSum4(nums []int, target int) int {
	res := make([]int, target+1)
	res[0] = 1
	for j := 0; j <= target; j++ { // 背包容量在外层
		for i := 0; i < len(nums); i++ { // 物品在里层
			if j-nums[i] >= 0 {
				res[j] += res[j-nums[i]]
			}
		}
		//log.Println(res)
	}

	return res[target]
}
```