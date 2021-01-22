/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) []int {
	for i := len(A) - 1; i >= 0; i-- {
		A[i] += K % 10
		K /= 10
		if A[i] > 9 {
			K++
			A[i] -= 10
		}
	}
	for K > 0 {
		A = append([]int{K % 10}, A...)
		K /= 10
	}
	return A
}
// @lc code=end

