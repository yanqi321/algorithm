/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (62.90%)
 * Likes:    534
 * Dislikes: 0
 * Total Accepted:    146.4K
 * Total Submissions: 232.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 
 * 
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
 * 
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
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root != nil {
		BFS([]*TreeNode{ root }, &ret)
	}
	return ret
}
func BFS(nodeList []*TreeNode, ret *[][]int) {
	if len(nodeList) == 0 {
		return
	}
	newList := []*TreeNode{}
	item := make([]int,len(nodeList))
	for i, v := range nodeList {
		item[i] = v.Val
		if v.Left != nil {
			newList = append(newList, v.Left)
		}
		if v.Right != nil {
			newList = append(newList, v.Right)
		}
	}
	*ret = append(*ret, item)
	BFS(newList, ret)
}
// @lc code=end

