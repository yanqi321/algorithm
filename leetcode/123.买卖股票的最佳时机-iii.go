/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxK := 2
	if len(prices) / 2 < maxK {
		maxK = len(prices) / 2
	}
	dptable := map[string]int{}
	for i := 0; i < len(prices); i++ {
		dptable[strconv.Itoa(i)+"-0-0"] = 0
		for k := maxK; k > 0; k-- {
			if i == 0 {
				dptable["0-"+strconv.Itoa(k)+"-0"] = 0
				dptable["0-"+strconv.Itoa(k)+"-1"] = -prices[i]
			} else {
				dptable[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-0"] = max(dptable[strconv.Itoa(i - 1)+"-"+strconv.Itoa(k)+"-0"], dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-1"] + prices[i])
				dptable[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-1"] = max(dptable[strconv.Itoa(i - 1)+"-"+strconv.Itoa(k)+"-1"], dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k-1)+"-0"] - prices[i])
			}
		}
	}
	return dptable[strconv.Itoa(len(prices) - 1)+"-"+strconv.Itoa(maxK) + "-0"]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

