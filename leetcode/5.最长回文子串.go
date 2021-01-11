/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// 暴力循环: 思路: 记录 startIndex 长度(不是最长字符串)
/**
* 状态 是否是回文串
* 状态方程  dp[i][j] = dp[i] == dp[j] && dp[i+1][j-1]
* start: dp[i][i] = true
* 边界  i + 1 < j - 1; j - i > 2;
*/
/**
* 中心扩散
**/
func longestPalindrome(s string) string {
	maxLen := 1
	ret := string(s[0])
	for i :=0 ; i < len(s) - 1; i++ {
		oddSp := centerSpread(s, i, i)
		evenSp := centerSpread(s, i, i +1)
		maxStr := oddSp
		if len(oddSp) < len(evenSp) {
			maxStr = evenSp
		}
		if len(maxStr) > maxLen {
			maxLen = len(maxStr)
			ret = maxStr
		}
	}
	return ret
}

func centerSpread(s string, left ,right int) string {
	for left >=0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return s[left + 1:right]
}
// @lc code=end

