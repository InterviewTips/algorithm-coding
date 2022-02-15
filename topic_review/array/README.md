### 数组专题复习

- 螺旋矩阵 II 记住左闭右开

loop,offset,startX,startY, 的应用

遍历长度 n + startX - offset

- 最短长度 连续子数组

双指针，left, right 开始位置都是 0, right 放 for 循环，left 移动的时机在 sum >= target 的时候 进行 `sum-=nums[left]` 并且 `left++`

- 有序平方数组

首尾指针，比较对应的平方，设立一个 index 从末尾开始，谁比较大则设置谁并移动指针，边界条件 left <= right