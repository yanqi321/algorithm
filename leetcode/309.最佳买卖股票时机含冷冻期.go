/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	prev_0 := 0
	dp_i_0 := max(prices[1]- prices[0], 0)
	dp_i_1 := max(-prices[0], -prices[1])
	for i := 2; i < len(prices); i++ {
		tmp0 := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
		dp_i_1 = max(dp_i_1, prev_0 - prices[i])
		prev_0 = tmp0
	}
	return dp_i_0
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

