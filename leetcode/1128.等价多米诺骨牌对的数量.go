/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	mp := make([]int,100) // 0-99
	ret := 0
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		val := v[0] * 10 + v[1]
		ret += mp[val] // 这个太巧妙了 n(n-1)/2 = n + n-1 + n-2 ... 
		mp[val]++
	}
	return ret
}
// @lc code=end

