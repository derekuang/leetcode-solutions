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
	ans := make([]int, len(A))
	even, odd := 0, 1
	for _, a := range A {
		if a%2 == 0 {
			ans[even] = a
			even += 2
		} else {
			ans[odd] = a
			odd += 2
		}
	}
	return ans
}

// @lc code=end
