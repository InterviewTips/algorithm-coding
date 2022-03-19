package leetcode

import (
	"encoding/json"
	"math"
)

var Null = math.MinInt64

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	// 层次遍历
	// 使用队列
	// 跟普通的层次遍历还不太一样 会加入 null 的打印
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		goto end
	}

	queue = append(queue, root)
	res = append(res, root.Val)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				res = append(res, node.Left.Val)
			} else {
				res = append(res, Null)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				res = append(res, node.Right.Val)
			} else {
				res = append(res, Null)
			}
		}
	}

	// 去除末尾的 null
	for res[len(res)-1] == Null {
		res = res[:len(res)-1]
	}

end:
	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(res string) *TreeNode {
	var data []int
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		panic(err)
	}

	// 构建
	if len(data) == 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: data[0]}
	queue = append(queue, head)
	index := 1 // 从下标索引 1 开始
	for index < len(data) {
		tmp := queue[0]
		queue = queue[1:]
		if index < len(data) && data[index] != Null { // -> left
			node := &TreeNode{Val: data[index]}
			tmp.Left = node
			queue = append(queue, node)
		}
		index++
		if index < len(data) && data[index] != Null { // -> right
			node := &TreeNode{Val: data[index]}
			tmp.Right = node
			queue = append(queue, node)
		}
		index++
	}

	return head
}
