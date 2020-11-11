/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

package solution

import "sort"

// @lc code=start
func sortArrayByParityII(A []int) []int {
	sort.Ints(A)
	for i := 0; i < len(A); {
		if (i%2 == 0 && A[i]%2 != 0) || (i%2 != 0 && A[i]%2 == 0) {
			remain := i % 2
			j := i + 1
			for A[j]%2 != remain {
				j++
			}
			t := A[j]
			for j > i {
				A[j] = A[j-1]
				j--
			}
			A[i] = t
			i += 2
		} else {
			i++
		}
	}
	return A
}

// @lc code=end
