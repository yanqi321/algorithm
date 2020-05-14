/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
// @lc code=start
import "unicode"
func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j {
		vi, vj := rune(s[i]), rune(s[j])
		if !isValid(vi) && !isValid(vj) {
			i++
			j--
		} else if !isValid(vi) {
			i++
		} else if !isValid(vj) {
			j--
		} else if unicode.ToUpper(vi) != unicode.ToUpper(vj) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
func isValid(v rune) bool {
	return unicode.IsDigit(v) || unicode.IsLetter(v)
}
// @lc code=end

