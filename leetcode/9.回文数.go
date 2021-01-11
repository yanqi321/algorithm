/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
// @lc code=start
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i, j := 0, len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
// @lc code=end

