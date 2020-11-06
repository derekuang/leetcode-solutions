/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

package solution

// @lc code=start
func findMin153(nums []int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) / 2
		if nums[mid] > nums[tail] {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	return nums[tail]
}

// @lc code=end
