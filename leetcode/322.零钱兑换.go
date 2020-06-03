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
	max int
	amount int
	coins []int
}

func coinChange(coins []int, amount int) int {
	s := newSolution(coins, amount)
	s.DFS(0, len(coins) - 1, 0)
	return s.max
}
func newSolution(coins []int, amount int) *solution {
	sort.Ints(coins)
	return &solution{ -1, amount, coins }
}
func (s *solution) DFS(currVal ,index, currCount int) {
	if s.amount == currVal {
		if currCount < s.max || s.max < 0 {
			s.max = currCount
			return
		}
	} else if currVal > s.amount {
		return
	}
	for i := index; i >=0; i-- {
		if currVal + s.coins[i] > s.amount {
			continue
		}
		if s.max > 0 && (s.amount - currVal) / s.coins[i] > s.max - currCount {
			break
		}
		currVal += s.coins[i]
		s.DFS(currVal, i, currCount + 1)
		currVal -= s.coins[i]
	}
	return
}
// @lc code=end

