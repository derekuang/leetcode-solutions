/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

package solution

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	p, q := head, head.Next
	for q != nil && q.Next != nil {
		p.Next = q.Next
		p = q.Next
		q.Next = p.Next
		q = p.Next
	}
	p.Next = evenHead
	return head
}

// @lc code=end
