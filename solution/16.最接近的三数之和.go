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
		for p2, p3 := p1+1, n-1; p2 < p3; {
			t = nums[p1] + nums[p2] + nums[p3]
			if t == target {
				return t
			}
			d1, d2 := abs16(ans-target), abs16(t-target)
			if d2 < d1 {
				ans = t
			}
			if t > target {
				p3--
			} else {
				p2++
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
