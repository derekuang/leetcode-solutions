/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

package solution

import "sort"

// @lc code=start
func largestPerimeter(A []int) int {
	sort.Ints(A)
	n := len(A)
	a, b, c := n-3, n-2, n-1
	for a >= 0 {
		if A[a]+A[b] > A[c] {
			return A[a] + A[b] + A[c]
		}
		a--
		b--
		c--
	}
	return 0
}

// @lc code=end
