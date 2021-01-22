/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(A []int) []bool {
	ret := make([]bool, len(A))
	prefix := 0
	for i :=0; i < len(A); i++ {
		prefix = (A[i] + prefix << 1) % 5 // 解决数据长度问题
		if prefix == 0 {
			ret[i] = true
		} else {
			ret[i] = false
		}
	}
	return ret
}
// @lc code=end

