/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (49.58%)
 * Likes:    1095
 * Dislikes: 0
 * Total Accepted:    80.4K
 * Total Submissions: 160.3K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// @lc code=start
func trap(height []int) int {
	stack := []int{}
	res := 0
	for i,v := range height {
		for len(stack) > 0 && v > height[stack[len(stack) - 1]] {
			h := height[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				break
			}
			minHeight := v
			width := i - stack[len(stack) - 1] - 1
			if height[stack[len(stack) - 1] ] < v {
				minHeight = height[stack[len(stack) - 1] ]
			}
			res += width * (minHeight - h)
		} 
		stack = append(stack, i)
	}
	return res
}
// @lc code=end

