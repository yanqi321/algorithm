/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := []byte{} // 单调栈
	for i:=0; i < len(num);i++ {
		for len(stack) >0 && num[i] < stack[len(stack) - 1] && k > 0 {
			stack = stack[:len(stack) - 1]
			k--
		}
		stack = append(stack, num[i])
	}
	// fmt.Println(stack)
	stack = stack[:len(stack) - k]
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	ans := string(stack)
	if ans == "" {
		ans = "0"
	}
	return ans
}
// @lc code=end

