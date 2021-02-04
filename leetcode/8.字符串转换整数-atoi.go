/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	len := len(s)
	index := 0
	for index < len && s[index] == ' ' { // 去除前置空格
		index++
	}
	if index == len {
		return 0
	}
	sign := 1
	if s[index] == '+' {
		index++
	} else if s[index] == '-' {
		index++
		sign = -1
	}
	ans := 0
	for index < len {
		if s[index] < '0' || s[index]> '9' {
			break
		}
		if ans > math.MaxInt32 / 10 || (ans == math.MaxInt32 / 10 && int(s[index] - '0') > (math.MaxInt32 % 10)) { // 可能不是整除
			return math.MaxInt32
		}
		if ans < math.MinInt32 / 10 || (ans == math.MinInt32 / 10 && int(s[index] - '0') > -(math.MinInt32 % 10)) {
			return math.MinInt32
		}
		// fmt.Println(ans)
		ans = ans * 10 + sign * int(s[index] - '0')
		index++
	}
	
	return ans
}
// @lc code=end

