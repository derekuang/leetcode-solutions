/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

package solution

// @lc code=start
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p := head
	for p != nil {
		q := p.Next
		p.Next = dummy.Next
		dummy.Next = p
		p = q
	}
	return dummy.Next
}

// @lc code=end
