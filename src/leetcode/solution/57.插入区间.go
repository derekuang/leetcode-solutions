/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

package solution

// @lc code=start
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	p := 0
	flag := true
	for p < len(intervals) && flag {
		l, r := newInterval[0], newInterval[1]
		set := make([]int, 2)
		if l > intervals[p][0] {
			set[0] = intervals[p][0]
			if l > intervals[p][1] {
				set[1] = intervals[p][1]
				p++
			} else {
				// l <= intervals[p][1]
				for p < len(intervals) && r > intervals[p][1] {
					p++
				}
				if p >= len(intervals) || r < intervals[p][0] {
					set[1] = r
				} else {
					set[1] = intervals[p][1]
					p++
				}
				flag = false
			}
		} else {
			// l <= intervals[p][0]
			set[0] = l
			if r < intervals[p][0] {
				set[1] = r
			} else {
				// r >= intervals[p][0]
				for p < len(intervals) && r > intervals[p][1] {
					p++
				}
				if p >= len(intervals) || r < intervals[p][0] {
					set[1] = r
				} else {
					set[1] = intervals[p][1]
					p++
				}
			}
			flag = false
		}
		ans = append(ans, set)
	}
	if flag {
		ans = append(ans, append([]int{}, newInterval...))
	}
	for p < len(intervals) {
		ans = append(ans, append([]int{}, intervals[p]...))
		p++
	}
	return
}

// @lc code=end
