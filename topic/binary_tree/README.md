<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [二叉树](#%E4%BA%8C%E5%8F%89%E6%A0%91)
  - [致谢](#%E8%87%B4%E8%B0%A2)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 二叉树

- 关于 leetcode 定义全局变量指针出现问题的解答

```go
// 如果定义了全局变量 测试样例并不会每次执行完就重置清空 所以有可能存在问题
var prev *TreeNode
func getMinimumDifference(root *TreeNode) int {
    prev = nil // 或者就是每次先在入口重置一下
```

#### 致谢

- [二叉树构造](https://blog.csdn.net/USTCsunyue/article/details/106148317)