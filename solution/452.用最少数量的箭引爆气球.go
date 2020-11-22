/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

package solution

import (
	"sort"
)

// @lc code=start
func findMinArrowShots(points [][]int) (count int) {
	n := len(points)
	if n < 2 {
		return n
	}

	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	p := n - 1
	for p >= 0 {
		count++
		target := points[p] // target points to the last balloon
		for p--; p >= 0 && target[0] <= points[p][1]; p-- {
		}
	}
	return
}

// @lc code=end
