/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// zeroIndex := 0
	// for i,v := range nums {
	// 	if (v != 0) {
	// 		nums[zeroIndex],nums[i] = v, nums[zeroIndex]
	// 		zeroIndex++
	// 	}
	// }
	for i, v := range nums {
		if v == 0 {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
				nums[j] = 0
			}
		}
	}
}
// @lc code=end

