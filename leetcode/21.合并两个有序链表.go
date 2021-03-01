/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := &ListNode{}
	cursor := headNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cursor.Next = l2
			break
		} else if l2 == nil {
			cursor.Next = l1
			break
		} else if l1.Val < l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			
			l2 = l2.Next
		}
		cursor = cursor.Next
	}
	return headNode.Next
}
// @lc code=end

