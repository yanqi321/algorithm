/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	w := 0
	for i :=0; i < len(nums); i++ {
		if w == i {
			continue
		}
		if (nums[i - 1] != nums[i]) { //  w 或者 i + 1 都可以
			w++
			nums[w] = nums[i]
		}
	}
	return w + 1
}
// @lc code=end

122344567
