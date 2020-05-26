/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (62.82%)
 * Likes:    591
 * Dislikes: 0
 * Total Accepted:    166.3K
 * Total Submissions: 262.9K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 * 
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [3,2,3]
 * 输出: 3
 * 
 * 示例 2:
 * 
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 * 
 * 
 */

// @lc code=start
func majorityElement(nums []int) int {
	res := trackBack(nums, 0, len(nums) - 1)
	return res
}
func trackBack(nums []int,l, r int) int {
	if l == r {
		return nums[l]
	}
	mid := (r - l) / 2 + l
	left := trackBack(nums, l, mid)
	right := trackBack(nums, mid + 1, r)

	if left == right {
		return left
	}

	countLeft := counter(nums, left, l, r)
	countRight := counter(nums, right, l, r)
	if countLeft > countRight {
		return left
	}
	return right
}
func counter(nums []int, n, l , r int) int {
	count := 0
	for ;l <= r; l++ {
		if nums[l] == n {
			count++
		}
	}
	return count
}
// @lc code=end

