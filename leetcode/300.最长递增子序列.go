/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	dp[0] = 1
	for i:=1; i< len(nums);i++ {
		maxDP :=0
		for j :=0; j < i; j++ {
			if nums[i] > nums[j] && maxDP < dp[j] {
				maxDP = dp[j]
			}
		}
		dp[i] = maxDP + 1
	}
	ret := 0
	for _,v := range dp {
		if v > ret {
			ret = v
		}
	}
	return ret
}
// @lc code=end

