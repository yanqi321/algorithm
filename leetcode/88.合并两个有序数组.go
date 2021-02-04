/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	index := m + n - 1
    i, j :=m - 1, n - 1;
    for index >= 0 {
        if i < 0 {
            nums1[index] = nums2[j]
            j--
        } else if j < 0 || nums1[i] > nums2[j] {
            nums1[index] = nums1[i]
            i--
        } else  {
            nums1[index] = nums2[j]
            j--
        }
        index--
    }
}
// @lc code=end

