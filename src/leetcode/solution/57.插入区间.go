/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

package solution

// @lc code=start
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	l, r := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > r {
			if !merged {
				ans = append(ans, []int{l, r})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < l {
			ans = append(ans, interval)
		} else {
			l = min(l, interval[0])
			r = max(r, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{l, r})
	}
	return
}

// @lc code=end
