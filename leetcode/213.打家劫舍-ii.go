/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	// 从第一间
	dp_i0 := nums[0]
	dp_i1 := max(nums[0], nums[1])
	if len(nums) == 2 {
		return dp_i1
	}
	// f(i) = max(f(i-1), f(i-2) + i)
	for i :=2; i < len(nums) - 1; i ++ {
		tmp := dp_i1
		dp_i1 = max(dp_i1, dp_i0 + nums[i])
		dp_i0 = tmp
	}
	leftMax := dp_i1
	// 从第二间
	dp_i0 = nums[1]
	dp_i1 = max(nums[1], nums[2])
	// f(i) = max(f(i-1), f(i-2) + i)
	for i :=3; i < len(nums); i ++ {
		tmp := dp_i1
		dp_i1 = max(dp_i1, dp_i0 + nums[i])
		dp_i0 = tmp
	}
	return max(leftMax, dp_i1)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

