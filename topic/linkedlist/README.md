<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [链表专题](#%E9%93%BE%E8%A1%A8%E4%B8%93%E9%A2%98)
  - [移除链表元素](#%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 链表专题

#### 移除链表元素

| 最近提交时间 | 题目                                                         | 题目难度 | 提交次数 |
| :----------- | :----------------------------------------------------------- | :------- | :------- |
| 1 分钟前     | [#203 移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/) | 简单     | 6 次     |

简单总结：

- 使用虚拟头节点 `dummyHead.Next = head`
- cur 变动 `cur.Next = cur.Next.Next`

