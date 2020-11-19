/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

package solution

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	t, ans := 0, nums[0]+nums[1]+nums[2]
	for p1 := 0; p1 < n-2; p1++ {
		for p2 := p1 + 1; p2 < n-1; p2++ {
			for p3 := p2 + 1; p3 < n; p3++ {
				t = nums[p1] + nums[p2] + nums[p3]
				d1, d2 := abs16(ans-target), abs16(t-target)
				if d2 < d1 {
					ans = t
				} else if d1 < d2 && t > ans {
					break
				}
			}
		}
	}
	return ans
}

func abs16(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
