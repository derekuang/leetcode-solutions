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

	l, m, r := 0, 0, mid
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target && (m == 0 || nums[m-1] != target) {
			break
		}
		if nums[m] == target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	low = m

	l, m, r = mid, 0, n-1
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target && (m == n-1 || nums[m+1] != target) {
			break
		}
		if nums[m] == target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	high = m

	return []int{low, high}
}

// @lc code=end
