/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (41.30%)
 * Likes:    1488
 * Dislikes: 0
 * Total Accepted:    243.8K
 * Total Submissions: 590.1K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 *range会隐式的unicode解码
 *除开rune和byte底层的类型的区别，在使用上，rune能处理一切的字符，而byte仅仅局限在ascii
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	keyMap := map[rune]rune {
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, v := range s {
		prev := len(stack) - 1
		if v == ')' || v == ']' || v == '}' {
			if prev >= 0 && keyMap[stack[prev]] == v {
				stack = stack[:len(stack) - 1]
				continue
			}
		}
		stack = append(stack, v)
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
// @lc code=end

