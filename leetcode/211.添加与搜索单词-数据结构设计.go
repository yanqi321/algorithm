/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type WordDictionary struct {
	isEnd bool
	next [26] *WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	for _, v := range word {
		if this.next[v - 'a'] == nil {
			this.next[v - 'a'] = &WordDictionary{}
		}
		this = this.next[v - 'a']
	}
  this.isEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(this, word)
}

func match (wordNode *WordDictionary, word string) bool {
	if len(wordNode.next) == 0 {
		return false
	}
	if len(word) == 1 {
		if word == '.' &&
	}
	if word[0] == '.' {
		for i := 0; i < 26; i++ {
			if match(wordNode.next[i], word[1:]) {
				return true
			}
		}
	}
	if wordNode.next[word[0] -'a'] != nil && match(wordNode.next[word[0] -'a'] , word[1:]) {
		return true
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

