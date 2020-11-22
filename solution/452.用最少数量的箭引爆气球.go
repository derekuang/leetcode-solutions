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
type arr452 [][]int

func (pt arr452) Len() int {
	return len(pt)
}
func (pt arr452) Less(i, j int) bool {
	if pt[i][0] < pt[j][0] {
		return true
	} else if pt[i][0] == pt[j][0] && pt[i][1] <= pt[j][1] {
		return true
	}
	return false
}
func (pt arr452) Swap(i, j int) {
	pt[i], pt[j] = pt[j], pt[i]
}

func findMinArrowShots(points [][]int) (count int) {
	n := len(points)
	if n < 2 {
		return n
	}

	sort.Sort(arr452(points))
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
