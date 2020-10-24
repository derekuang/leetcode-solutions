/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

package solution

// @lc code=start

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	abs := func(x, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}

	lHeight := height(root.Left)
	if lHeight == -1 {
		return -1
	}
	rHeight := height(root.Right)
	if rHeight == -1 {
		return -1
	}
	if abs(lHeight, rHeight) > 1 {
		return -1
	}

	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}

// @lc code=end
