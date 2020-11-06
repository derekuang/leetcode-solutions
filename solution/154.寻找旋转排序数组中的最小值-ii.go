/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

package solution

// @lc code=start
func findMin(nums []int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) / 2
		if nums[mid] < nums[tail] {
			tail = mid
		} else if nums[mid] > nums[tail] {
			head = mid + 1
		} else {
			tail--
		}
	}
	return nums[head]
}

// @lc code=end
