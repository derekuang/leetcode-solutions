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
	hash1, hash2 := map[int]int{}, map[int]int{}
	for _, n0 := range uniques[0] {
		c0 := counters[0][n0]
		for _, n1 := range uniques[1] {
			c1 := counters[1][n1]
			hash1[n0+n1] += c0 * c1
		}
	}
	for _, n2 := range uniques[2] {
		c2 := counters[2][n2]
		for _, n3 := range uniques[3] {
			c3 := counters[3][n3]
			hash2[n2+n3] += c2 * c3
		}
	}

	for k, v1 := range hash1 {
		v2 := hash2[-k]
		if v2 > 0 {
			ans += v1 * v2
		}
	}
	return
}

// @lc code=end
