package _146

type LinkedList struct {
	Key  int
	Val  int
	Prev *LinkedList
	Next *LinkedList
}

type LRUCache struct {
	count    int
	capacity int
	cache    map[int]*LinkedList
	head     *LinkedList
	tail     *LinkedList
}

func Constructor(capacity int) LRUCache {
	head := &LinkedList{}
	head.Prev = nil
	tail := &LinkedList{}
	tail.Next = nil
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LinkedList),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.cache[key]
	if ok { // 更新
		l.moveToHead(v)
		return v.Val
	} else {
		return -1
	}
}

func (l *LRUCache) removeNode(node *LinkedList) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (l *LRUCache) moveToHead(node *LinkedList) {
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) Put(key int, value int) {
	v, ok := l.cache[key]
	if !ok {
		node := &LinkedList{
			Key:  key,
			Val:  value,
			Prev: nil,
			Next: nil,
		}
		l.cache[key] = node
		l.addNode(node)
		l.count++
		if l.count > l.capacity {
			tail := l.popTail()
			delete(l.cache, tail.Key)
			l.count--
		}
		return
	} else {
		v.Val = value
		l.moveToHead(v)
	}
}

func (l *LRUCache) popTail() *LinkedList {
	tail := l.tail.Prev
	l.removeNode(tail)
	return tail
}

func (l *LRUCache) addNode(node *LinkedList) {
	node.Prev = l.head
	node.Next = l.head.Next

	l.head.Next.Prev = node
	l.head.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
