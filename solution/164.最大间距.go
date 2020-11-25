/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

package solution

import "sort"

// @lc code=start
func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	sort.Ints(nums)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > ans {
			ans = nums[i] - nums[i-1]
		}
	}
	return
}

// @lc code=end
