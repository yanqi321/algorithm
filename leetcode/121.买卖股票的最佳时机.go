/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 状态机 动态规划
	arrLen := len(prices)
	if arrLen == 0 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, -prices[0]
	for i :=1; i< arrLen; i++ {
		// dpMap[strconv.Itoa(i) + "-0"] = max(dpMap[strconv.Itoa(i - 1) + "-0"], prices[i] + dpMap[strconv.Itoa(i - 1) + "-1"])
		dp_i_0 = max(dp_i_0, prices[i] + dp_i_1)
		// dpMap[strconv.Itoa(i) + "-1"] = max(dpMap[strconv.Itoa(i - 1) + "-1"], -prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

