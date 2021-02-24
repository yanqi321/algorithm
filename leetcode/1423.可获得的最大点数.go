/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
// 方法一 逆向思维 滑动窗口 前缀和(求剩余窗口最小值)
func maxScore(cardPoints []int, k int) int {
	wLen := len(cardPoints) - k
	for i := 1; i < len(cardPoints); i++ {
			cardPoints[i] += cardPoints[i - 1]
	}
	if wLen < 1 {  // 需要考虑窗口为0的特殊情况
			return cardPoints[len(cardPoints) - 1]
	}
	minVal := 0
	for i := wLen; i < len(cardPoints); i++ {
			newVal := cardPoints[i] - cardPoints[i - wLen]
			if newVal < minVal {
					minVal = newVal
			}
	}
	return cardPoints[len(cardPoints) - 1] - minVal
}
//方法2 dfs 超时了。。。
func maxScore2(cardPoints []int, k int) int {
	dp := &dpTable{
			dt: make(map[string]int),
			cardPoints: cardPoints,
	}
	ret := dp.dfs(0, len(cardPoints) - 1, k)
	return ret
}
type dpTable struct {
	dt map[string]int
	cardPoints []int
}
func (dp *dpTable) dfs(s, e, k int) int {
	key := getKey(s, e, k)
	if dp.dt[key] != 0 || s > e {
			return dp.dt[key]
	}
	if k == 1 {
			if dp.cardPoints[s] > dp.cardPoints[e] {
					dp.dt[key] = dp.cardPoints[s]
			} else {
					dp.dt[key] = dp.cardPoints[e]
			}
			return dp.dt[key]
	}
	// 取左边
	sVal := dp.cardPoints[s] + dp.dfs(s + 1, e, k-1)
	// 取右边
	eVal := dp.cardPoints[e] + dp.dfs(s, e - 1, k-1)
	if sVal > eVal {
			dp.dt[key] = sVal
	} else {
			dp.dt[key] = eVal
	}
	return dp.dt[key]
}
func getKey(s, e, k int) string {
	return strconv.Itoa(s) + "-" + strconv.Itoa(e) + "-" + strconv.Itoa(k)
}
// @lc code=end

