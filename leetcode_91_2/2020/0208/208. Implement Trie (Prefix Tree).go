package _208

type Node struct {
	isWord bool
	next   map[byte]*Node
}

func NewNode(isWord bool) *Node {
	return &Node{isWord: isWord, next: make(map[byte]*Node)}
}

type Trie struct {
	root *Node
}

//func NewTrie() *Trie {
//	return &Trie{root: NewNode(false)}
//}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: NewNode(false)}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		_, ok := cur.next[c]
		if !ok { //
			cur.next[c] = NewNode(false)
		}
		cur = cur.next[c]
	}
	if !cur.isWord { // 标记为单词
		cur.isWord = true
	}
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		v, ok := cur.next[c]
		if !ok {
			return false
		}
		cur = v
	}
	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		v, ok := cur.next[c]
		if !ok {
			return false
		}
		cur = v
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
