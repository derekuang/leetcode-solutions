/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

package solution

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	var mergeCount func(arr []int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}

		arr1 := append([]int{}, arr[:n/2]...)
		arr2 := append([]int{}, arr[n/2:]...)
		cnt := mergeCount(arr1) + mergeCount(arr2)

		l, r := 0, 0
		for _, v := range arr1 {
			for l < len(arr2) && arr2[l]-v < lower {
				l++
			}
			for r < len(arr2) && arr2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		p, q := 0, 0
		for i := 0; i < len(arr); i++ {
			if p < len(arr1) && (q == len(arr2) || arr1[p] <= arr2[q]) {
				arr[i] = arr1[p]
				p++
			} else {
				arr[i] = arr2[q]
				q++
			}
		}

		return cnt
	}

	prefixSums := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSums[i+1] = prefixSums[i] + v
	}
	return mergeCount(prefixSums)
}

// @lc code=end
