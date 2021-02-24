/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	nodeList :=[]*TreeNode{ root }
	ret := root.Val
	for len(nodeList) > 0 {
		node := nodeList[0]
		ret = node.Val
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
		}
		nodeList = nodeList[1:]
	}
	return ret
}
// @lc code=end

