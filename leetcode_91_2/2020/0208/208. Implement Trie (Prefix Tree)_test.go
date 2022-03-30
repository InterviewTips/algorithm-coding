package _208

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // 返回 true
	println(trie.Search("app"))     // 返回 false
	println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	println(trie.Search("app")) // 返回 true
}
