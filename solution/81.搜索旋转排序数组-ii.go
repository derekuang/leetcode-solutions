/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

package solution

// @lc code=start
func search81(nums []int, target int) bool {
	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := head + (tail-head)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < nums[tail] {
			if target > nums[mid] && target <= nums[tail] {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		} else if nums[mid] > nums[tail] {
			if target < nums[mid] && target > nums[tail] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		} else {
			tail--
		}
	}

	return false
}

// @lc code=end
