/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

package solution

// @lc code=start
func searchInsert(nums []int, target int) int {
	head, mid, tail := 0, 0, len(nums)-1
	for head <= tail {
		mid = (head + tail) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}

	if target < nums[mid] {
		return mid
	}
	return mid + 1
}

// @lc code=end
