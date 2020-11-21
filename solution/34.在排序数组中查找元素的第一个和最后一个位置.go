/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

package solution

// @lc code=start
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	low, mid, high := 0, 0, n-1
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			break
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if nums[mid] != target {
		return []int{-1, -1}
	}
	for low = mid; low-1 >= 0 && nums[low-1] == target; low-- {
	}
	for high = mid; high+1 < n && nums[high+1] == target; high++ {
	}
	return []int{low, high}
}

// @lc code=end
