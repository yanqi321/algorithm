/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (39.24%)
 * Likes:    509
 * Dislikes: 0
 * Total Accepted:    37.3K
 * Total Submissions: 95.2K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	stack := make([]int, 1)
	stack[0] = -1
	maxArea := 0
	for  i, v := range heights {
		for len(stack) > 1 && v < heights[stack[len(stack) - 1]] {
			index := stack[len(stack) - 1]
			currArea := heights[index] * (i - stack[len(stack) - 2] - 1)
			stack = stack[:len(stack) - 1]
			if currArea > maxArea {
				maxArea = currArea
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		currArea := heights[stack[len(stack)-1]] * (len(heights) - stack[len(stack) - 2] - 1)
		stack = stack[:len(stack) - 1]
		if currArea > maxArea {
			maxArea = currArea
		}
	}
	return maxArea
}
// @lc code=end

