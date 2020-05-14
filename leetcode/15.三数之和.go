/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (26.26%)
 * Likes:    1956
 * Dislikes: 0
 * Total Accepted:    188.7K
 * Total Submissions: 718.6K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例：
 * 
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */
// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i >0 &&nums[i] == nums[i-1] {
			continue
		}
		l,r := i + 1, len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{ nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				} 
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				r--
				l++
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
// @lc code=end

