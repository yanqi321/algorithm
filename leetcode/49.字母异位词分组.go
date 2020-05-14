/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (61.19%)
 * Likes:    335
 * Dislikes: 0
 * Total Accepted:    71.6K
 * Total Submissions: 116.7K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	hashMap := make(map[string][]string)
	for _,str := range strs {
		key := getKey(str)
		hashMap[key] = append(hashMap[key], str)
	}
	for _,v := range hashMap {
		res = append(res, v)
	}
	return res
}

func getKey(str string) string {
	arr:=[]rune(str)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]>arr[j]
	})
	return string(arr)
}
// @lc code=end

