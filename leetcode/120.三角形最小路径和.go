/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (64.15%)
 * Likes:    405
 * Dislikes: 0
 * Total Accepted:    60.5K
 * Total Submissions: 93.4K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 * 
 * 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
 * 
 * 
 * 
 * 例如，给定三角形：
 * 
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 * 
 * 
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 * 
 * 
 * 
 * 说明：
 * 
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 * 
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	dpMap := map[string]int{}
	ret := dp(0, 0, triangle, dpMap)
	return ret
}
func dp(i, j int, triangle [][]int, dpMap map[string]int) int {
	if i + 1 == len(triangle) {
		return triangle[i][j]
	}
	key1 := strconv.Itoa(i+1) + "#" + strconv.Itoa(j)
	key2 := strconv.Itoa(i+1) + "#" + strconv.Itoa(j+1)
	if _,ok := dpMap[key1]; !ok {
		dpMap[key1] = dp(i+1,j, triangle, dpMap)
	}
	if _,ok := dpMap[key2]; !ok {
		dpMap[key2] = dp(i+1,j + 1, triangle, dpMap)
	}
	return min(dpMap[key1],dpMap[key2]) + triangle[i][j]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

