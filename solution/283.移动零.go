/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

package solution

// @lc code=start
func moveZeroes(nums []int) {
	n, p, q := len(nums), 0, 0
	for q < n {
		if nums[q] != 0 {
			nums[p], nums[q] = nums[q], nums[p]
			p++
		}
		q++
	}
	return
}

// @lc code=end
