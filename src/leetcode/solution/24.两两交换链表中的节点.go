/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

package solution

// ListNode Data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head
	p, q := head, head.Next
	// start swap one pair
	p.Next = q.Next
	q.Next = dummy.Next
	dummy.Next = q

	p.Next = swapPairs(p.Next)

	return dummy.Next
}

// @lc code=end
