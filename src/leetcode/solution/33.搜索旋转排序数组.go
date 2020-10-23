/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

package solution

// @lc code=start
func search33(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := head + (tail-head)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[tail] {
			if target > nums[mid] && target <= nums[tail] {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		} else {
			if target < nums[mid] && target > nums[tail] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		}
	}

	return -1
}

// @lc code=end
