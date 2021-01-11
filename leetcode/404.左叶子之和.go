/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	ret := 0
	if root == nil {
		return ret
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		treeNode := stack[0]
		stack = stack[1:]
		if treeNode.Left != nil && treeNode.Left.Left == nil && treeNode.Left.Right == nil {
			ret += treeNode.Left.Val
		} else if treeNode.Left != nil {
			stack = append(stack, treeNode.Left)
		}
		if treeNode.Right != nil{
			stack = append(stack, treeNode.Right)
		}
	}
	return ret
}
// @lc code=end

