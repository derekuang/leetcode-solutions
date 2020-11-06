/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

package solution

// @lc code=start
func search704(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := (head + tail) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}

	return -1
}

// @lc code=end
