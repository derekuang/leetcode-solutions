/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

package solution

import "math"

// @lc code=start

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lDepth := getDepth(root.Left)
	rDepth := getDepth(root.Right)
	if math.Abs(float64(lDepth-rDepth)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := getDepth(root.Left)
	rDepth := getDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// @lc code=end
