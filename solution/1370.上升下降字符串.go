/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

package solution

import "sort"

// @lc code=start
func sortString(s string) (ans string) {
	sLen := len(s)
	if sLen < 2 {
		return s
	}

	counter := map[int]int{}
	for _, c := range s {
		counter[int(c)]++
	}
	unique := []int{}
	for c := range counter {
		unique = append(unique, c)
	}
	sort.Ints(unique)

	cnt := 0
	appendAns := func(i int) {
		used := map[int]bool{}
		c := unique[i]
		if !used[c] && counter[c] > 0 {
			used[c] = true
			counter[c]--
			cnt++
			ans = ans + string(rune(c))
		}
	}

	flag := -1 // -1: find min;	1: find max
	for cnt < sLen {
		if flag == -1 {
			for i := 0; i < len(unique); i++ {
				appendAns(i)
			}
		} else {
			for i := len(unique) - 1; i >= 0; i-- {
				appendAns(i)
			}
		}
		flag *= -1
	}
	return
}

// @lc code=end
