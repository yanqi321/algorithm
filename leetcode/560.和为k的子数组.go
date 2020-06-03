/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.49%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    51.4K
 * Total Submissions: 115.7K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 * 
 * 示例 1 :
 * 
 * 
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 * 
 * 
 * 说明 :
 * 
 * 
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 * 
 * 
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	preSum,count := 0, 0
	preMap := map[int]int{}
	for _,v := range nums {
		preSum += v
		if preSum	== k {
			count++
		}
		filterSum := preSum - k
		if cnt, ok := preMap[filterSum]; ok {
			count += cnt
		}
		if _, ok := preMap[preSum]; !ok {
			preMap[preSum] = 1
		} else {
			preMap[preSum]++
		}
	}
	return count
}
// @lc code=end

