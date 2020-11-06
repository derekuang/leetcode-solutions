/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

package solution

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	d := make([]int, len(nums)+1)
	d[1] = nums[0]

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[res] {
			res++
			d[res] = nums[i]
		} else {
			// find the number lesser than nums[i]
			l, r, pos := 1, res, 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}

	return res
}

// @lc code=end
