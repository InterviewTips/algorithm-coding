<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [数组专题](#%E6%95%B0%E7%BB%84%E4%B8%93%E9%A2%98)
  - [二分查找](#%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE)

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

