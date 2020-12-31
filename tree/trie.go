package main

import "fmt"

type Trie struct {
	isEnd bool
	next [26]*Trie
	val int
}

func (t *Trie) Insert(key string, val int) {
	for _, v := range key {
		if t.next[v - 'a'] == nil {
			t.next[v - 'a'] = new(Trie)
		}
		t = t.next[v - 'a']
	}
	t.isEnd = true
	t.val = val
}

func (t *Trie) Contain(key string) bool {
	for _, v := range key {
		if t.next[v - 'a'] == nil {
			return false
		}
		t = t.next[v - 'a']
	}
	return t.isEnd
}

func (t *Trie) PreSearch(key string) int {
	ret := 0
	for _, v := range key {
		if t.next[v - 'a'] == nil {
			return ret
		}
		t = t.next[v - 'a']
	}
	return ret + bfs(t)
}

func bfs(t *Trie) int {
	ret := 0
	if t.isEnd {
		ret += t.val
	}
	stack := make([]*Trie,0)
	stack = append(stack, t)
	for len(stack) > 0 {
		node := stack[0]
		if node == nil {
			stack = stack[1:]
			continue
		}
		if node.isEnd {
			ret += node.val
		}
		for _, v := range node.next {
			stack = append(stack, v)
		}
		stack = stack[1:]
	}
	return ret
}

func NewTrie() Trie {
	return Trie{}
}

func main() {
	trie := NewTrie()
	trie.Insert("apple", 2)
	contain := trie.Contain("apple")
	contain2 := trie.Contain("app")
	trie.Insert("applewatch", 4)
	trie.Insert("appgirl", 9)
	pre := trie.PreSearch("app")
	fmt.Println(contain, contain2, pre)
}