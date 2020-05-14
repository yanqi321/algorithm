/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (72.73%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    22.1K
 * Total Submissions: 30.1K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的后序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其后序遍历: [5,6,3,2,4,1].
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
		res := make([]int, 0)
		postSort(root, &res)
		return res
}
func postSort(node *Node, res *[]int) {
	if node != nil {
		if len(node.Children) != 0 {
			for _,n := range node.Children {
				postSort(n, res)
			}
		}
		*res = append(*res, node.Val)
	}
}
// @lc code=end

