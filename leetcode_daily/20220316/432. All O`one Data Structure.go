package _0220316

import "container/list"

type node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	*list.List                          // 内置数据结构 双向链表
	nodes      map[string]*list.Element // 存储每个 key 对应的链表节点
}

func Constructor() AllOne {
	return AllOne{
		List:  list.New(), // 初始化双向链表
		nodes: map[string]*list.Element{},
	}
}

func (a *AllOne) Inc(key string) {
	if cur := a.nodes[key]; cur != nil { // key 在链表中
		curNode := cur.Value.(node) // 断言
		// 查询当前节点的下一个节点，如果为空，或者下一个节点的 count 大于当前节点的 count+1, 则在当前节点后新建一个节点存 count+1
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
			a.nodes[key] = a.InsertAfter(
				node{
					keys: map[string]struct{}{
						key: {},
					},
					count: curNode.count + 1,
				},
				cur,
			)
		} else { // 正常，那么加入 key 到下一个节点
			nxt.Value.(node).keys[key] = struct{}{}
			a.nodes[key] = nxt // 此时 key 对应的节点为 nxt
		}
		// 需要将 key 从当前节点删除
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 { // 如果当前节点没有 key 了 则删除
			a.Remove(cur)
		}
	} else { // key 不在链表中 则新增一个节点
		if a.Front() == nil || a.Front().Value.(node).count > 1 {
			a.nodes[key] = a.PushFront(node{map[string]struct{}{key: {}}, 1})
		} else { // 不为空且 count == 1
			a.Front().Value.(node).keys[key] = struct{}{}
			a.nodes[key] = a.Front()
		}
	}
}

func (a *AllOne) Dec(key string) {
	// 测试用例已经假定一定存在 key 对应的节点
	cur := a.nodes[key]
	curNode := cur.Value.(node)
	if curNode.count > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(node).count < curNode.count-1 { // 前置节点不符合条件 新建
			a.nodes[key] = a.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
		} else { // 符合条件
			pre.Value.(node).keys[key] = struct{}{}
			a.nodes[key] = pre
		}
	} else { // key 仅出现一次，将其移出 nodes
		delete(a.nodes, key)
	}
	// 删除在当前节点的 key
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		a.Remove(cur)
	}
}

func (a *AllOne) GetMaxKey() string {
	if b := a.Back(); b != nil { // 取最后面的 key
		for key := range b.Value.(node).keys { // 随机返回
			return key
		}
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	if f := a.Front(); f != nil { // 取最前面的 key
		for key := range f.Value.(node).keys { // 随机返回
			return key
		}
	}
	return ""
}
