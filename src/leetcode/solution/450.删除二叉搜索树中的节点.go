/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 */

package solution

// @lc code=start
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	dummy := &TreeNode{}
	dummy.Left = root

	pre, target := dummy, root
	for target != nil && target.Val != key {
		pre = target
		if target.Val > key {
			target = target.Left
		} else {
			target = target.Right
		}
	}
	if target == nil {
		return root
	}

	if target.Left == nil {
		if pre.Left == target {
			pre.Left = target.Right
		} else {
			pre.Right = target.Right
		}
	} else if target.Right == nil {
		if pre.Left == target {
			pre.Left = target.Left
		} else {
			pre.Right = target.Left
		}
	} else {
		l, r := target.Left, target.Right
		if pre.Left == target {
			pre.Left = l
		} else {
			pre.Right = l
		}
		for l.Right != nil {
			l = l.Right
		}
		l.Right = r
	}
	return dummy.Left
}

// @lc code=end
