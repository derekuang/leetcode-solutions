/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

package solution

// TreeNode Data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type anchor struct {
	start, end int
}

func generateTrees(n int) []*TreeNode {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// Calculate the number of binary search trees can be generated with n
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		acc := 0
		for j := 1; j <= i; j++ {
			l := max(1, nums[j-1])
			r := max(1, nums[i-j])
			acc += (l * r)
		}
		nums[i] = acc
	}
	// Generate binary search trees recursively
	treeMap := make(map[anchor][]*TreeNode)
	res := generate(1, n, nums, treeMap)
	return res
}

func generate(start, end int, nums []int, treeMap map[anchor][]*TreeNode) []*TreeNode {
	if start > end {
		return nil
	}
	if treeMap[anchor{start, end}] != nil {
		return treeMap[anchor{start, end}]
	}
	size := end - start + 1
	index, res := 0, make([]*TreeNode, nums[size])
	// Generate tree lists that uses i as the head node
	for i := start; i <= end; i++ {
		lLists := generate(start, i-1, nums, treeMap)
		rLists := generate(i+1, end, nums, treeMap)
		if lLists == nil {
			lLists = []*TreeNode{nil}
		}
		if rLists == nil {
			rLists = []*TreeNode{nil}
		}
		for _, left := range lLists {
			for _, right := range rLists {
				head := new(TreeNode)
				head.Val = i
				head.Left = left
				head.Right = right
				res[index] = head
				index++
			}
		}
	}
	treeMap[anchor{start, end}] = res

	return res
}

// @lc code=end
