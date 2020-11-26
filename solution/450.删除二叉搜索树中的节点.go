/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 */

package solution

// @lc code=start
func deleteNode450(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode450(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode450(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = root.Right
		root = root.Left
	}

	return root
}

// @lc code=end
