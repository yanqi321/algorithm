/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (53.02%)
 * Likes:    707
 * Dislikes: 0
 * Total Accepted:    111.7K
 * Total Submissions: 208.9K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 * 
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 * 
 */

// @lc code=start
func letterCombinations(digits string) []string {
	digMap := map[byte][]string{
		'2': { "a", "b", "c" },
		'3': { "d","e", "f" },
		'4': { "g", "h", "i" },
		'5': { "j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p","q","r","s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := []string{}
	if len(digits) > 0 {
		backTrick(digits,"", &res, digMap)
	}
	return res
}

func backTrick(digits string, tmp string, res *[]string, digMap map[byte][]string) {
	if len(digits) == 0 {
		*res = append(*res, tmp)
		return
	}
	curr := digits[0]
	strs := digMap[curr]
	for _, v := range strs {
		newTmp := tmp + v
		backTrick(digits[1:], newTmp, res, digMap)
	}
}
// @lc code=end

