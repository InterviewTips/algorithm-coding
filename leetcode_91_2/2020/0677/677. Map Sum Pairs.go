package _677

type Node struct {
	isWord bool
	val    int
	next   map[byte]*Node
}

func NewNode() *Node {
	return &Node{next: make(map[byte]*Node)}
}

type MapSum struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{root: NewNode()}
}

func (m *MapSum) Insert(key string, val int) {
	cur := m.root
	for i := 0; i < len(key); i++ {
		c := key[i]
		_, ok := cur.next[c]
		if !ok {
			cur.next[c] = NewNode()
		}
		cur = cur.next[c]
	}
	cur.val = val // 赋值
	cur.isWord = true
}

func (m *MapSum) Sum(prefix string) int {
	cur := m.StartsWith(prefix)
	//递归搜寻节点
	return m.sum(cur)
}

func (m *MapSum) sum(root *Node) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.isWord {
		res += root.val
	}
	for _, v := range root.next {
		res += m.sum(v)
	}

	return res
}

// StartsWith 找到符合前缀的最后一个节点
func (m *MapSum) StartsWith(prefix string) *Node {
	cur := m.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		v, ok := cur.next[c]
		if !ok {
			return nil
		}
		cur = v
	}
	return cur
}
