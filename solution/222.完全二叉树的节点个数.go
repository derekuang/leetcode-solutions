/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 */

package solution

// @lc code=start
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 1
	stack := []*TreeNode{root}
	for true {
		node := stack[0]
		stack = stack[1:]
		if node.Left != nil {
			if node.Right != nil {
				count += 2
				stack = append(stack, node.Left, node.Right)
			} else {
				count++
				break
			}
		} else {
			break
		}
	}
	return count
}

// @lc code=end
