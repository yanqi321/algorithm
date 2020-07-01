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
	index := 0
	ret := calHelper(s, &index)
	return ret
}

func calHelper (s string, i *int) int {
	stack := make([]int,0)
	sgin := '+'
	num := 0
	for ;*i < len(s);*i++ {
		if  unicode.IsDigit(rune(s[*i])) {
			num = 10 * num + (int(s[*i]) - '0')
		}
		if s[*i] == '(' {
			*i++
			num = calHelper(s, i)
			*i++
		}
		if (!unicode.IsDigit(rune(s[*i])) && rune(s[*i]) != ' ') || len(s) <= *i + 1 {
			if sgin == '+' {
				stack = append(stack, num)
			} else if sgin == '-' {
				stack = append(stack, -num)
			} else if sgin == '*' {
				pre := stack[len(stack) - 1]
				stack[len(stack) - 1] = pre * num
			} else if sgin == '/' {
				pre := stack[len(stack) - 1]
				stack[len(stack) - 1] = pre / num
			}
			if s[*i] == ')' {
				break
			}
			num = 0
			sgin = rune(s[*i])
		}
	}
	ret := 0
	for _,v := range stack {
		ret += v
	}
	return ret
}
// @lc code=end

