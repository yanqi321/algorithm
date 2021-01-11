package main

import "fmt"

type trieNode struct {
	isEnd bool
	val int
	next map[rune]*trieNode
}

func (t *trieNode) AddWord(word string, val int) {
	node := t
	for _, v := range word {
		if node.next[v] == nil {
			node.next[v] = &trieNode{false, 0, map[rune]*trieNode{}}
		}
		node = node.next[v]
	}
	node.isEnd = true
	node.val = val
}

func (t *trieNode) Contain(word string) bool {
	node := t
	for _, v := range word {
		if node.next[v] == nil {
			return false
		}
		node = node.next[v]
	}
	return node.isEnd
}

func (t *trieNode) PreSearch(key string) int {
	node := t
	for _, v := range key {
		if node.next[v] == nil {
			return 0
		}
		node = node.next[v]
	}
	return dfs(t)
}

func dfs(node *trieNode) int {
	val := 0
	if node == nil {
		return val
	}
	if node.isEnd {
		val += node.val
	}
	if node.next == nil {
		return val
	}
	for _, v := range node.next {
		val += dfs(v)
	}
	return val
}

func Constructor() *trieNode {
	return &trieNode{false, 0, map[rune]*trieNode{}}
}


func main() {
	// trie := Constructor()
	// trie.AddWord("apple", 2)
	// contain := trie.Contain("apple")
	// contain2 := trie.Contain("app")
	// trie.AddWord("applewatch", 4)
	// trie.AddWord("appgirl", 9)
	// pre := trie.PreSearch("app")
	// fmt.Println(contain, contain2, pre)
	mm := map[int]string{}
	mm[1] = "a"
	fmt.Println("test", mm[1])
}