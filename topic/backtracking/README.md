<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [回溯](#%E5%9B%9E%E6%BA%AF)
- [todo](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 回溯

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 13 分钟前    | [#51 N 皇后](https://leetcode-cn.com/problems/n-queens/)     | 困难     | 2 次     |
| 31 分钟前    | [#37 解数独](https://leetcode-cn.com/problems/sudoku-solver/) | 困难     | 3 次     |
| 1 天前       | [#47 全排列 II](https://leetcode-cn.com/problems/permutations-ii/) | 中等     | 1 次     |
| 1 天前       | [#46 全排列](https://leetcode-cn.com/problems/permutations/) | 中等     | 2 次     |
| 1 天前       | [#491 递增子序列](https://leetcode-cn.com/problems/increasing-subsequences/) | 中等     | 4 次     |
| 2 天前       | [#90 子集 II](https://leetcode-cn.com/problems/subsets-ii/)  | 中等     | 2 次     |
| 2 天前       | [#78 子集](https://leetcode-cn.com/problems/subsets/)        | 中等     | 1 次     |
| 2 天前       | [#93 复原 IP 地址](https://leetcode-cn.com/problems/restore-ip-addresses/) | 中等     | 3 次     |
| 2 天前       | [#131 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/) | 中等     | 2 次     |
| 2 天前       | [#40 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/) | 中等     | 2 次     |
| 2 天前       | [#39 组合总和](https://leetcode-cn.com/problems/combination-sum/) | 中等     | 4 次     |
| 2 天前       | [#17 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/) | 中等     | 4 次     |
| 2 天前       | [#216 组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii/) | 中等     | 3 次     |
| 2 天前       | [#77 组合](https://leetcode-cn.com/problems/combinations/)   | 中等     | 3 次     |

回溯模版

```
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
```

### todo

- [ ] [#332 重新安排行程](https://leetcode-cn.com/problems/reconstruct-itinerary/)