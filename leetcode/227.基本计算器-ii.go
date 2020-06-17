/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (36.20%)
 * Likes:    139
 * Dislikes: 0
 * Total Accepted:    17.8K
 * Total Submissions: 49K
 * Testcase Example:  '"3+2*2"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 * 
 * 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
 * 
 * 示例 1:
 * 
 * 输入: "3+2*2"
 * 输出: 7
 * 
 * 
 * 示例 2:
 * 
 * 输入: " 3/2 "
 * 输出: 1
 * 
 * 示例 3:
 * 
 * 输入: " 3+5 / 2 "
 * 输出: 5
 * 
 * 
 * 说明：
 * 
 * 
 * 你可以假设所给定的表达式都是有效的。
 * 请不要使用内置的库函数 eval。
 * 
 * 
 */

// @lc code=start
func calculate(s string) int {
	stack = make([]int, 0)
	flag = '+'
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if v == '+' {
			num = 0
			flag = '+'
		} else if v == '-' {
			num = 0
			flag = '-'
		} else if v == ' ' {
			num = 0
		}  else if v == '*' {
			pre := stack[len(stack) - 1]
			
		} else if v == '/' {
			num = 0
		}else {
			num = 10 * num + strconv.Atoi(v)
		}
	}
}
// @lc code=end

