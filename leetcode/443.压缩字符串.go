/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	w, anchor := 0, 0
	for r := 0; r < len(chars); r++ {
		if r + 1 == len(chars) || chars[r + 1] != chars[r] {
			chars[w] = chars[anchor];
			w++
			if (r > anchor) {
				num := strconv.Itoa(r + 1 - anchor)
				for _, val := range num {
					chars[w] = byte(val)
					w++
				}
			}
			anchor = r + 1
		}
	}
	return w
}
// @lc code=end

