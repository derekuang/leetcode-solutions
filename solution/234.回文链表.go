/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

package solution

// @lc code=start
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = reverLinedList234(slow.Next)
	l, mid, r := head, slow.Next, slow.Next
	for l != mid {
		if l.Val != r.Val {
			return false
		}
		l, r = l.Next, r.Next
	}
	if r == nil || r.Next == nil {
		return true
	}
	return false
}

func reverLinedList234(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	for head != nil {
		t := dummy.Next
		dummy.Next = head
		head = head.Next
		dummy.Next.Next = t
	}
	return dummy.Next
}

// @lc code=end
