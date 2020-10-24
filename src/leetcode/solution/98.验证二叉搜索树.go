/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	// Push all the mostleft elements
	top, stack := 0, []*TreeNode{}
	for p := root; p != nil; p = p.Left {
		stack = append(stack, p)
		top++
	}

	var pre *TreeNode = nil
	for len(stack) > 0 {
		// Pop out the top element
		top--
		cur := stack[top]
		stack = stack[0:top]
		// The value of cur node should be greater than the pre's
		if pre != nil && cur.Val <= pre.Val {
			return false
		}
		pre = cur
		if cur.Right != nil {
			for p := cur.Right; p != nil; p = p.Left {
				stack = append(stack, p)
				top++
			}
		}
	}

	return true
}

// @lc code=end

