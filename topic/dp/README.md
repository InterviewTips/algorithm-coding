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

- 做题记录

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 5 小时前     | [#516 最长回文子序列](https://leetcode-cn.com/problems/longest-palindromic-subsequence/) | 中等     | 1 次     |
| 6 小时前     | [#647 回文子串](https://leetcode-cn.com/problems/palindromic-substrings/) | 中等     | 1 次     |
| 6 小时前     | [#72 编辑距离](https://leetcode-cn.com/problems/edit-distance/) | 困难     | 1 次     |
| 7 小时前     | [#583 两个字符串的删除操作](https://leetcode-cn.com/problems/delete-operation-for-two-strings/) | 中等     | 3 次     |
| 7 小时前     | [#115 不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/) | 困难     | 1 次     |
| 8 小时前     | [#392 判断子序列](https://leetcode-cn.com/problems/is-subsequence/) | 简单     | 2 次     |
| 9 小时前     | [#53 最大子数组和](https://leetcode-cn.com/problems/maximum-subarray/) | 简单     | 8 次     |
| 9 小时前     | [#1035 不相交的线](https://leetcode-cn.com/problems/uncrossed-lines/) | 中等     | 2 次     |
| 10 小时前    | [#1143 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/) | 中等     | 2 次     |
| 1 天前       | [#718 最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/) | 中等     | 1 次     |
| 1 天前       | [#674 最长连续递增序列](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/) | 简单     | 2 次     |
| 1 天前       | [#300 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/) | 中等     | 4 次     |
| 2 天前       | [#714 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | 中等     | 2 次     |
| 2 天前       | [#309 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | 中等     | 2 次     |
| 2 天前       | [#188 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/) | 困难     | 3 次     |
| 2 天前       | [#123 买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/) | 困难     | 2 次     |
| 2 天前       | [#122 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/) | 中等     | 6 次     |
| 2 天前       | [#121 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/) | 简单     | 2 次     |
| 3 天前       | [#337 打家劫舍 III](https://leetcode-cn.com/problems/house-robber-iii/) | 中等     | 4 次     |
| 3 天前       | [#213 打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/) | 中等     | 2 次     |
| 3 天前       | [#198 打家劫舍](https://leetcode-cn.com/problems/house-robber/) | 中等     | 2 次     |
| 3 天前       | [#139 单词拆分](https://leetcode-cn.com/problems/word-break/) | 中等     | 1 次     |
| 4 天前       | [#279 完全平方数](https://leetcode-cn.com/problems/perfect-squares/) | 中等     | 2 次     |
| 4 天前       | [#322 零钱兑换](https://leetcode-cn.com/problems/coin-change/) | 中等     | 3 次     |
| 4 天前       | [#70 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/) | 简单     | 5 次     |
| 4 天前       | [#377 组合总和 Ⅳ](https://leetcode-cn.com/problems/combination-sum-iv/) | 中等     | 1 次     |
| 5 天前       | [#518 零钱兑换 II](https://leetcode-cn.com/problems/coin-change-2/) | 中等     | 1 次     |
| 5 天前       | [#474 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/) | 中等     | 1 次     |
| 5 天前       | [#494 目标和](https://leetcode-cn.com/problems/target-sum/)  | 中等     | 1 次     |
| 6 天前       | [#1049 最后一块石头的重量 II](https://leetcode-cn.com/problems/last-stone-weight-ii/) | 中等     | 3 次     |
| 6 天前       | [#416 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/) | 中等     | 3 次     |
| 7 天前       | [#96 不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/) | 中等     | 1 次     |
| 7 天前       | [#343 整数拆分](https://leetcode-cn.com/problems/integer-break/) | 中等     | 5 次     |
| 7 天前       | [#62 不同路径](https://leetcode-cn.com/problems/unique-paths/) | 中等     | 2 次     |
| 7 天前       | [#63 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/) | 中等     | 3 次     |
| 7 天前       | [#746 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/) | 简单     | 2 次     |
| 8 天前       | [#509 斐波那契数](https://leetcode-cn.com/problems/fibonacci-number/) | 简单     | 2 次     |