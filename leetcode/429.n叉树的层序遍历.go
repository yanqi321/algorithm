/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (65.06%)
 * Likes:    81
 * Dislikes: 0
 * Total Accepted:    19.5K
 * Total Submissions: 29.8K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其层序遍历:
 * 
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 * 
 * 
 * 
 * 
 * 说明:
 * 
 * 
 * 树的深度不会超过 1000。
 * 树的节点总数不会超过 5000。
 * 
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
		res := make([][]int,0)
		if root !=nil {
			traversalNode(root, &res, 0)
		}
		return res
}
func traversalNode(node *Node, res *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], node.Val)
	for _, n := range node.Children {
		traversalNode(n, res, level+1)
	}
}
// @lc code=end

