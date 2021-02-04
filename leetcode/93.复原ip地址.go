/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
var result []string
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	result = []string{}
	track := make([]string, 0)
	backtrack(s, track, 0)
	return result
}

func backtrack(s string, track []string, segIdx int) {
	if segIdx == 4 {
		if len(s) == 0 {
			ipAddr := strings.Join(track, ".")
			result = append(result, ipAddr)
		}
	}
	if  (4 - segIdx) > len(s) || 3 * (4 - segIdx) < len(s) {
		// fmt.Printf("需要 %d 个片段，字符串长度 %d \n", 3 - segIdx, len(s))
		return
	}
	for i:=1; i <= 3; i++ {
		if i > len(s) {
			return
		} 
		v, err := strconv.Atoi(s[:i])
		// fmt.Printf("第 %d 段，选择 %d位 s: %s, 值%d \n", segIdx, i, s[:i], v)
		if err == nil && v < 256 {
			track = append(track, s[:i])
			str := s[i:]
			backtrack(str, track, segIdx + 1)
			track = track[:len(track) - 1]
		}
		if v == 0 {
			break
		}
	}
}


// @lc code=end

