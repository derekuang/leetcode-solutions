/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

package solution

import "sort"

// @lc code=start
func intersection(nums1 []int, nums2 []int) (ans []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if i > 0 && nums1[i] == nums1[i-1] {
			i++
			continue
		}

		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return
}

// @lc code=end
