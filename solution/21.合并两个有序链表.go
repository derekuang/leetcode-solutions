/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

package solution

// @lc code=start
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	d := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			d.Next = l1
			l1 = l1.Next
		} else {
			d.Next = l2
			l2 = l2.Next
		}
		d = d.Next
	}
	if l1 != nil {
		d.Next = l1
	} else {
		d.Next = l2
	}
	return dummy.Next
}

// @lc code=end
