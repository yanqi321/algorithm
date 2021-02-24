/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	mp := make(map[byte]int, 26)
	for i := range s {
		v := s[i]
		mp[v - 'a']++
	}
	stack := []byte{}
	for i := range s {
		v := s[i]
		if isContain(stack, v) {
			mp[v - 'a']--
			continue
		}
		for len(stack) > 0 && mp[stack[len(stack) - 1] - 'a'] > 1 && v < stack[len(stack) - 1] {
			mp[stack[len(stack) - 1] - 'a']--
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, v)
	}
	return string(stack)
}

func isContain (stack []byte, v byte) bool {
	for _, val := range stack {
		if v == val {
			return true
		}
	}
	return false
}
// @lc code=end

