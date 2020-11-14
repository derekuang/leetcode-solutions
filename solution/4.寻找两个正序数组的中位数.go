/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

package solution

// @lc code=start
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	mergeNums := make([]int, totalLen)
	p, q, m := 0, 0, 0
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] <= nums2[q] {
			mergeNums[m] = nums1[p]
			p++
		} else {
			mergeNums[m] = nums2[q]
			q++
		}
		m++
	}
	for p < len(nums1) {
		mergeNums[m] = nums1[p]
		p++
		m++
	}
	for q < len(nums2) {
		mergeNums[m] = nums2[q]
		q++
		m++
	}
	mid := totalLen / 2
	if totalLen%2 == 0 {
		return float64(mergeNums[mid]+mergeNums[mid-1]) / 2
	}
	return float64(mergeNums[mid])
}

// @lc code=end
