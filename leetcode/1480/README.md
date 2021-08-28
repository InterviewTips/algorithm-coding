<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [1480. 一维数组的动态和](#1480-%E4%B8%80%E7%BB%B4%E6%95%B0%E7%BB%84%E7%9A%84%E5%8A%A8%E6%80%81%E5%92%8C)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 1480. 一维数组的动态和

```
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。

请返回 nums 的动态和。

示例 1：

输入：nums = [1,2,3,4]
输出：[1,3,6,10]
解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
示例 2：

输入：nums = [1,1,1,1,1]
输出：[1,2,3,4,5]
解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。
示例 3：

输入：nums = [3,1,2,10,1]
输出：[3,4,6,16,17]

提示：

1 <= nums.length <= 1000
-10^6 <= nums[i] <= 10^6
```

https://leetcode-cn.com/problems/running-sum-of-1d-array