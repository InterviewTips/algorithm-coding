package _297

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/InterviewTips/algorithm-coding/guns"
)

type TreeNode = guns.TreeNode

type Codec struct {
	s string
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	c.middleList(root)
	return c.s[:len(c.s)-1] // 去除最后一个 ,
}

func (c *Codec) middleList(root *TreeNode) {
	if root == nil {
		c.s += "nil,"
		return
	}
	c.s += fmt.Sprintf("%d,", root.Val)
	c.middleList(root.Left)
	c.middleList(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	fmt.Printf("data is %s\n", data)
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		c.l = append(c.l, l[i])
	}
	return c.buildTree()
}

func (c *Codec) buildTree() *TreeNode {
	if c.l[0] == "nil" {
		c.l = c.l[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.l[0])
	root := &TreeNode{Val: val}
	c.l = c.l[1:]
	root.Left = c.buildTree()
	root.Right = c.buildTree()
	return root
}
