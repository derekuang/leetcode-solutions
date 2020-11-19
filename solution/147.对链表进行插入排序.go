/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

package solution

// @lc code=start
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	p, pre, q := dummy, head, head.Next
	for q != nil {
		for p != pre && p.Next.Val < q.Val {
			p = p.Next
		}
		if p != pre {
			pre.Next = q.Next
			q.Next = p.Next
			p.Next = q
		} else {
			pre = pre.Next
		}
		p = dummy
		q = pre.Next
	}
	return dummy.Next
}

// @lc code=end
