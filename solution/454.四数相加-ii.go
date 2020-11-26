/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

package solution

import "sort"

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) (ans int) {
	nums := [][]int{A, B, C, D}
	counters := [4]map[int]int{}
	for i := 0; i < 4; i++ {
		counters[i] = map[int]int{}
		for _, num := range nums[i] {
			counters[i][num]++
		}
	}
	uniques := [4][]int{}
	for i, m := range counters {
		t := []int{}
		for num := range m {
			t = append(t, num)
		}
		sort.Ints(t)
		uniques[i] = t
	}

	for i := 0; i < len(uniques[0]); i++ {
		n0 := uniques[0][i]
		c0 := counters[0][n0]
		for j := 0; j < len(uniques[1]); j++ {
			n1 := uniques[1][j]
			c1 := counters[1][n1]
			for k := 0; k < len(uniques[2]); k++ {
				n2 := uniques[2][k]
				c2 := counters[2][n2]

				rest := -n0 - n1 - n2
				c3 := counters[3][rest]
				if c3 > 0 {
					ans += c0 * c1 * c2 * c3
				}
			}
		}
	}
	return
}

// @lc code=end
