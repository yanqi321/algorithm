/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	arrLen := len(flowerbed)
	for i:=0; i < arrLen; i++ {
		if flowerbed[i] == 1 {
			i++ // 下格0
			continue
		} else {
			isLeft := true
			isRight := true
			if i > 0	&& flowerbed[i - 1] == 1 {
				isLeft = false
			}
			if i + 1 < arrLen && flowerbed[i + 1] == 1  {
			isRight = false
			}
			if isLeft && isRight {
				n--
				flowerbed[i] = 1
				i++
			}
		}
		if n <= 0 {
			return true
		}
	}
	return n == 0
}
// @lc code=end

