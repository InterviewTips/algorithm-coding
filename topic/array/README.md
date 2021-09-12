<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [数组专题](#%E6%95%B0%E7%BB%84%E4%B8%93%E9%A2%98)
  - [二分查找](#%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 数组专题

#### 二分查找

- 区间的重要性
- middle 计算要防止溢出 `(left + right)/2`  如果两个数字很大有可能会溢出，所以可以使用 `left + ((right - left) >> 2)`