/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.00%)
 * Likes:    564
 * Dislikes: 0
 * Total Accepted:    88K
 * Total Submissions: 114K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

// @lc code=start
// recursion 递归
/* func subsets(nums []int) [][]int {
	output := make([][]int, 0)
	output = append(output, []int{})
	for _, v := range nums {
		newList := [][]int{}
		for _, item := range output {
			newItem := make([]int, len(item))
			copy(newItem, item)
			newItem = append(newItem, v)
			newList = append(newList, newItem)
		}
		output = append(output, newList...)
	}
	return output
} */
func subsets(nums []int) [][]int {
	output :=make([][]int, 0)
	output = append(output, []int{})
	cur := make([]int, 0)
	trackback(nums, 0, &cur, &output)
	return output
}

func trackback(nums []int, i int, cur *[]int, output *[][]int) {
	if i == len(nums) {
		return
	}
	*cur = append(*cur, nums[i])
	newItem := make([]int,len(*cur));
	copy(newItem, *cur)
	*output = append(*output, newItem)
	trackback(nums, i + 1, cur, output)
	*cur = (*cur)[:len(*cur) - 1]
	trackback(nums, i+ 1, cur, output)
}
// @lc code=end

