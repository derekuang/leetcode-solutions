/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

package solution

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	p := 0
	for q := 1; q < n; q++ {
		for q < n && nums[q] == nums[q-1] {
			q++
		}
		if q < n {
			p++
			nums[p] = nums[q]
		}
	}
	return p + 1
}

// @lc code=end
