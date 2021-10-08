<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [数组专题](#%E6%95%B0%E7%BB%84%E4%B8%93%E9%A2%98)
  - [二分查找](#%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE)
  - [移除元素](#%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0)
  - [滑动窗口](#%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3)
  - [模拟行为(螺旋矩阵)](#%E6%A8%A1%E6%8B%9F%E8%A1%8C%E4%B8%BA%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 数组专题

#### 二分查找

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 44 分钟前    | [#367 有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/) | 简单     | 1 次     |
| 1 小时前     | [#69 x 的平方根](https://leetcode-cn.com/problems/sqrtx/)    | 简单     | 5 次     |
| 1 小时前     | [#34 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | 中等     | 2 次     |
| 2 小时前     | [#35 搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/) | 简单     | 3 次     |
| 2 小时前     | [#704 二分查找](https://leetcode-cn.com/problems/binary-search/) | 简单     | 1 次     |

简单总结：

- 区间的重要性(左闭右闭还是左闭右开)
- middle 计算要防止溢出 `(left + right)/2` 如果两个数字很大有可能会溢出，所以可以使用 `left + ((right - left) >> 2)`
- 如果是要查找大于等于 target 的第一个索引，则最后 return left (可参考 34、35 题)
- 如果是要查找小于等于 target 的最后一个索引，则最后 return right (可参考 69 题)

#### 移除元素

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 2 分钟前     | [#977 有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/) | 简单     | 2 次     |
| 1 天前       | [#844 比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare/) | 简单     | 2 次     |
| 2 天前       | [#283 移动零](https://leetcode-cn.com/problems/move-zeroes/) | 简单     | 1 次     |
| 2 天前       | [#26 删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) | 简单     | 6 次     |
| 2 天前       | [#27 移除元素](https://leetcode-cn.com/problems/remove-element/) | 简单     | 3 次     |

简单总结：

- 双指针的基本思路 多练习

#### 滑动窗口

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 5 分钟前     | [#76 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/) | 困难     | 2 次     |
| 2 小时前     | [#904 水果成篮](https://leetcode-cn.com/problems/fruit-into-baskets/) | 中等     | 4 次     |
| 3 天前       | [#209 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/) | 中等     | 3 次     |

简单总结:

- 双指针+哈希表

#### 模拟行为(螺旋矩阵)

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 37 分钟前    | [#剑指 Offer 29 顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/) | 简单     | 2 次     |
| 44 分钟前    | [#54 螺旋矩阵](https://leetcode-cn.com/problems/spiral-matrix/) | 中等     | 3 次     |
| 2 小时前     | [#59 螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/) | 中等     | 1 次     |
