/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode-cn.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (44.38%)
 * Likes:    287
 * Dislikes: 0
 * Total Accepted:    30.7K
 * Total Submissions: 69.3K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或
 * -中选择一个符号添加在前面。
 * 
 * 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：nums: [1, 1, 1, 1, 1], S: 3
 * 输出：5
 * 解释：
 * 
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 * 
 * 一共有5种方法让最终目标和为3。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 数组非空，且长度不会超过 20 。
 * 初始的数组的和不会超过 1000 。
 * 保证返回的最终结果能被 32 位整数存下。
 * 
 * 
 */

// @lc code=start
type s struct {
	nums []int
	dpTable map[string]int
}
func newWays(nums []int) *s {
	return &s{
		nums, map[string]int{},
	}
}
func findTargetSumWays(nums []int, S int) int {
	s := newWays(nums)
	return s.dp(len(nums) - 1, S)
}
func (s *s) dp(index, val int) int {
	if index == 0 {
		if val == 0 && s.nums[index] == 0 {
			return 2
		} else if val == s.nums[index] || -val == s.nums[index] {
			return 1
		}
		return 0
	}
	key := strconv.Itoa(index-1) + "#" + strconv.Itoa(val)
	if _,ok := s.dpTable[key]; !ok {
		s.dpTable[key] = s.dp(index - 1, val + s.nums[index]) + s.dp(index-1, val - s.nums[index])
	}
	return s.dpTable[key]
}
// @lc code=end

