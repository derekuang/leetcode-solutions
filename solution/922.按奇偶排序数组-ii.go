/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

package solution

// @lc code=start
func sortArrayByParityII(A []int) []int {
	for even, odd := 0, 1; even < len(A); even += 2 {
		if A[even]%2 != 0 {
			for A[odd]%2 != 0 {
				odd += 2
			}
			A[even], A[odd] = A[odd], A[even]
		}
	}
	return A
}

// @lc code=end
