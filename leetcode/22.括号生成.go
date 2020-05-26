/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (73.76%)
 * Likes:    1019
 * Dislikes: 0
 * Total Accepted:    126.5K
 * Total Submissions: 167.9K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 * 
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {
	output := make([]string, 0)
	trackBack(n, n, "",  &output)
	return output
}

func trackBack(l int, r int, cur string, output *[]string) {
	if l ==0 && r == 0 {
		*output = append(*output, cur)
		return
	}
	if l > 0 {
		cur += "("
		trackBack(l -1, r, cur, output)
		cur = cur[:len(cur) - 1]
	}
	if r > l {
		cur += ")"
		trackBack(l, r - 1, cur, output)
	}
}
// @lc code=end

