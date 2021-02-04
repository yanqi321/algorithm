/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
/*
* 
*/
func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	setA :=map[int]struct{}{}
	for _, v := range A {
		sumA += v
		setA[v] = struct{}{}
	}
	sumB := 0 
	for _, v := range B {
		sumB += v
	}

	delta := (sumA + sumB) / 2
	for _, y := range B {
		x := delta - (sumB - y)
		if _, has := setA[x]; has {
			return []int{x, y}
		}
	}
	return []int{}
}
// @lc code=end

