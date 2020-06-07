/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	if len(prices) / 2 < k {
		k = len(prices)/ 2
	}
	dpTable := make([][2]int, k + 1)
	for i,v := range prices {
		dpTable[0] = [2]int{ 0 }
		for j :=1; j <= k; j++ {
			if i == 0 {
				dpTable[j] = [2]int{ 0, -v }
			} else {
				dpTable[j][0] = max(dpTable[j][0],dpTable[j][1] + v)
				dpTable[j][1] = max(dpTable[j][1],dpTable[j-1][0] - v)
			}
		}
	}
	return dpTable[k][0]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end
