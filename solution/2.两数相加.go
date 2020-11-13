/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

package solution

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	p, p1, p2, ad := &ListNode{}, l1, l2, 0
	ans = new(ListNode)
	p.Next = ans
	for p1 != nil || p2 != nil || ad != 0 {
		var n1, n2 int
		if p1 != nil {
			n1, p1 = p1.Val, p1.Next
		}
		if p2 != nil {
			n2, p2 = p2.Val, p2.Next
		}
		sum := n1 + n2 + ad
		ad, sum = sum/10, sum%10
		p.Next.Val, p.Next.Next = sum, new(ListNode)
		p = p.Next
	}
	p.Next = nil
	return
}

// @lc code=end
