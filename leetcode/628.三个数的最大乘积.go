/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */
import "sort"
// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	len := len(nums)
	maxVal1 := nums[len - 3] * nums[len - 2] * nums[ len - 1]
	maxVal2 := nums[0] * nums[1] * nums[len-1]
	if maxVal1 > maxVal2 {
		return maxVal1
	}
	return maxVal2
}
// @lc code=end

