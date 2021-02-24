/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */

// @lc code=start
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := []int{}
	len1, len2 := len(nums1), len(nums2)
	for i := 0; i <= k;i++ {
		if k - i > len2 {
			continue
		}
		if i > len1 {
			break
		}
		a := maxSub(nums1, len1 -i)
		b := maxSub(nums2, len2 - k + i)
		cans := merge(a, b)
		// fmt.Println(i, a, b)
		if isLess(ans, cans) {
			ans = cans
		}
	}
	return ans
}

func merge(a, b []int) (arr []int) {
	arr = make([]int, len(a) + len(b))
	for i := range arr {
		if isLess(a, b) {
			arr[i], b = b[0],b[1:]
		} else {
			arr[i], a = a[0], a[1:] 
		}
	}
	return
}

func isLess(a, b []int) bool {
	for i :=0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
func maxSub(nums []int, k int) []int {
	stack := []int{}
	for i := range nums {
		for len(stack) > 0 && stack[len(stack) - 1] < nums[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, nums[i])
	}
	stack = stack[:len(stack) - k]
	return stack
}
// @lc code=end

