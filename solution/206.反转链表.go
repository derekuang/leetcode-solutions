/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

package solution

// @lc code=start
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// @lc code=end
