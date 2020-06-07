/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (38.78%)
 * Likes:    606
 * Dislikes: 0
 * Total Accepted:    87.8K
 * Total Submissions: 221.2K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3 
 * 解释: 11 = 5 + 5 + 1
 * 
 * 示例 2:
 * 
 * 输入: coins = [2], amount = 3
 * 输出: -1
 * 
 * 
 * 
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 * 
 */

// @lc code=start
type solution struct {
	amount int
	coins []int
	dpMap map[int]int
}

func coinChange(coins []int, amount int) int {
	s := newSolution(coins, amount)
	return s.dp(amount)
}
func newSolution(coins []int, amount int) *solution {
	// sort.Ints(coins)
	return &solution{ amount, coins, map[int]int{} }
}
func (s *solution) dp(currVal int) int {
	if currVal == 0 {
		return 0
	} else if currVal < 0 {
		return -1
	}
	if v, ok := s.dpMap[currVal]; ok {
		return v
	}
	min := -2
	for _, v := range s.coins {
		res := s.dp(currVal - v)
		if res >= 0 && (res < min|| min == -2) {
			min = res 
		}
	}
	s.dpMap[currVal] = min + 1
	return min + 1
}
// @lc code=end

