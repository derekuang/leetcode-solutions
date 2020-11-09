/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

package solution

import "sort"

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	p, q := -1, 0
	for tp, tq := 0, 1; tq < len(nums); {
		if nums[tp] < nums[tq] {
			p, q = tp, tq
		} else if nums[tp] > nums[tq] {
			if p >= 0 && nums[tq] > nums[p] && nums[tq] < nums[q] {
				q = tq
			}
		}
		tp++
		tq++
	}
	if p >= 0 {
		nums[p], nums[q] = nums[q], nums[p]
	}
	sort.Ints(nums[p+1:])
}

// @lc code=end
