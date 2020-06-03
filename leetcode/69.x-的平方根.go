/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.63%)
 * Likes:    407
 * Dislikes: 0
 * Total Accepted:    157.2K
 * Total Submissions: 405.9K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
	l, r, ret:= 0, x, 0
	for l <= r {
		half := l + (r - l) / 2
		if half * half <= x {
			ret = half
			l = half + 1
		} else {
			r = half - 1
		}
	}
	return ret
}
// @lc code=end

