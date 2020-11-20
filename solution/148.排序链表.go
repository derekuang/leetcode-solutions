/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

package solution

// @lc code=start
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	len := 0
	for p := head; p != nil; p = p.Next {
		len++
	}

	dummy := &ListNode{Next: head}
	for subLen := 1; subLen < len; subLen <<= 1 {
		pre, cur := dummy, dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var q *ListNode
			if cur != nil {
				q = cur.Next
				cur.Next = nil
			}
			cur = q

			pre.Next = merge148(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return dummy.Next
}

func merge148(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	d := dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			d.Next = head1
			head1 = head1.Next
		} else {
			d.Next = head2
			head2 = head2.Next
		}
		d = d.Next
	}
	if head1 != nil {
		d.Next = head1
	}
	if head2 != nil {
		d.Next = head2
	}
	return dummy.Next
}

// @lc code=end
