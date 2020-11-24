/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

package solution

// @lc code=start
func intersect(nums1 []int, nums2 []int) (ans []int) {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if m[num] > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}
	return
}

// @lc code=end
