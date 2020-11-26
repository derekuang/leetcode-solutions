/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

package solution

// @lc code=start
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}

// @lc code=end
