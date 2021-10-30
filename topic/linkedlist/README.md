<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [链表专题](#%E9%93%BE%E8%A1%A8%E4%B8%93%E9%A2%98)
  - [移除链表元素](#%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 链表专题

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 2 分钟前     | [#142 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/) | 中等     | 1 次     |
| 5 天前       | [#面试题 02.07 链表相交](https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/) | 简单     | 2 次     |
| 6 天前       | [#19 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/) | 中等     | 4 次     |
| 9 天前       | [#24 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) | 中等     | 3 次     |
| 10 天前      | [#206 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/) | 简单     | 6 次     |
| 13 天前      | [#707 设计链表](https://leetcode-cn.com/problems/design-linked-list/) | 中等     | 5 次     |
| 21 天前      | [#203 移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/) | 简单     | 6 次     |

#### 移除链表元素

简单总结：

- 使用虚拟头节点 `dummyHead.Next = head`
- cur 变动 `cur.Next = cur.Next.Next`
