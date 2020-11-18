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
		for p < n-1 && nums[p] != 0 {
			p++
		}
		if p >= n-1 {
			return
		}
		if p+1 > q {
			q = p + 1
		}
		for q < n && nums[q] == 0 {
			q++
		}
		if q >= n {
			return
		}
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q++
	}
	return
}

// @lc code=end
