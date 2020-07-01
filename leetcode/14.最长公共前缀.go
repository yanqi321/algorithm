/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longPre := strs[0]
	for i :=1; i < len(strs); i++ {
		str := strs[i]
		strLen := min(len(longPre), len(str))
		j :=0
		for ; j < strLen; j++ {
			if longPre[j] != str[j] {
				break
			}
		}
		longPre = longPre[:j]
	}
	return longPre
}
func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}
// @lc code=end

