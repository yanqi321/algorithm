/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (56.39%)
 * Likes:    453
 * Dislikes: 0
 * Total Accepted:    64.3K
 * Total Submissions: 114.1K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 * 
 * 示例 1:
 * 
 * 输入: n = 12
 * 输出: 3 
 * 解释: 12 = 4 + 4 + 4.
 * 
 * 示例 2:
 * 
 * 输入: n = 13
 * 输出: 2
 * 解释: 13 = 4 + 9.
 * 
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+ 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j :=1; j * j <= i; j++ {
			dp[i] = min(dp[i], dp[i - j * j] + 1)
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
// @lc code=end

