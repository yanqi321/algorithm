/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (62.27%)
 * Likes:    1244
 * Dislikes: 0
 * Total Accepted:    167.3K
 * Total Submissions: 268.5K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 * 
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 * 
 * 
 * 
 * 
 * 
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 
 */

// @lc code=start
func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	maxArea := (j - i) * min(height[i],height[j])
	for i < j{
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
		newArea := (j - i) * min(height[i],height[j])
		if newArea > maxArea {
			maxArea = newArea
		}
	}
	return maxArea
}
func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
// @lc code=end

