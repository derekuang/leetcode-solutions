/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

package solution

import "sort"

// @lc code=start
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nums := []int{}
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	sort.Ints(nums)

	p := head
	for _, num := range nums {
		p.Val = num
		p = p.Next
	}

	return head
}

// @lc code=end
