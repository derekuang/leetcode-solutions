/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 */

package solution

// @lc code=start

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}

// @lc code=end
