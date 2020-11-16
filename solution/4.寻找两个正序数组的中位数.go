/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

package solution

import "math"

// @lc code=start
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	var pos1, pos2 int
	if totalLen%2 == 0 {
		pos1, pos2 = totalLen/2-1, totalLen/2
	} else {
		pos1, pos2 = totalLen/2, totalLen/2
	}
	acc, p, q := 0, 0, 0
	target1, target2 := 0, 0
	nums1 = append(nums1, math.MaxInt64)
	nums2 = append(nums2, math.MaxInt64)
	for acc <= pos2 {
		if nums1[p] <= nums2[q] {
			if acc == pos1 {
				target1 = nums1[p]
			}
			if acc == pos2 {
				target2 = nums1[p]
			}
			p++
			acc++
		} else {
			if acc == pos1 {
				target1 = nums2[q]
			}
			if acc == pos2 {
				target2 = nums2[q]
			}
			q++
			acc++
		}
	}
	return float64(target1+target2) / 2
}

// @lc code=end
