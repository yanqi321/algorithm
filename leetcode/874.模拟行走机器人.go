/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 *
 * https://leetcode-cn.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (33.44%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 23.2K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * 机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
 * 
 * 
 * -2：向左转 90 度
 * -1：向右转 90 度
 * 1 <= x <= 9：向前移动 x 个单位长度
 * 
 * 
 * 在网格上有一些格子被视为障碍物。
 * 
 * 第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])
 * 
 * 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。
 * 
 * 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入: commands = [4,-1,3], obstacles = []
 * 输出: 25
 * 解释: 机器人将会到达 (3, 4)
 * 
 * 
 * 示例 2：
 * 
 * 输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * 输出: 65
 * 解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * 答案保证小于 2 ^ 31
 * 
 * 
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	dx := []int{0, -1, 0, 1}
	dy := []int{1, 0, -1, 0}
	di, maxArea, endX, endY := 0 , 0, 0 ,0
	blockMap := make(map[string]bool, len(obstacles))
	for _, v := range obstacles {
		key := strconv.Itoa(v[0]) + "#" + strconv.Itoa(v[1])
		blockMap[key] = true
	}
	for _, v := range commands {
		if v == -1 {
			di = (di + 3) % 4
		} else if v == -2 {
			di = (di + 1) % 4
		} else {
			for i := 0; i < v; i++ {
				nextX := endX + dx[di]
				nextY := endY + dy[di]
				key := strconv.Itoa(nextX) + "#" + strconv.Itoa(nextY)
				if _, ok := blockMap[key]; ok {
					break
				}
				endX, endY = nextX, nextY
				maxArea = max(maxArea, endX * endX + endY * endY)
			}
		}
	}
	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} 
// @lc code=end

