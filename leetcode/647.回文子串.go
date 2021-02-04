/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	sLen := len(s)
	ret := 0
	for i := 0; i < sLen; i++ {
		odd := centerSpread(s, i, i+ 1)
		even := centerSpread(s, i, i)
		ret += odd + even
	}
	return ret
}

func centerSpread(s string, i, j int) (count int) {
	for i >=0 && j < len(s) {
		if s[i] != s[j] {
			return
		}
		count++
		i--
		j++
	}
	return
}
// @lc code=end

