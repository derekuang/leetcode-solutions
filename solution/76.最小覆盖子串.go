/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

package solution

import "strings"

// @lc code=start
func minWindow(s string, t string) string {
	cnts, needs := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}

	s = trim(s, t)

	isContain := func() bool {
		for k, v := range needs {
			if cnts[k] < v {
				return false
			}
		}
		return true
	}

	start, end := -1, len(s)
	for l, r := 0, 0; r < len(s); r++ {
		cnts[s[r]]++

		for isContain() && l <= r {
			if r-l < end-start {
				start, end = l, r
			}
			cnts[s[l]]--
			l++
		}
	}

	if start == -1 {
		return ""
	}
	return s[start : end+1]
}

func trim(s, t string) string {
	lBorder, rBorder := 0, len(s)
	for lBorder < rBorder && !strings.ContainsRune(t, rune(s[lBorder])) {
		lBorder++
	}
	for rBorder > lBorder && !strings.ContainsRune(t, rune(s[rBorder-1])) {
		rBorder--
	}
	return s[lBorder:rBorder]
}

// @lc code=end
