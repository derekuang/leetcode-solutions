/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

package solution

import "strings"

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	needs := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		needs[s1[i]]++
	}

	// Return state code
	contain := func(cnts map[byte]int) int {
		for k, v := range needs {
			if cnts[k] < v {
				return 1
			} else if cnts[k] > v {
				return 2
			}
		}
		return 0
	}

	for l, r := 0, 0; r < len(s2); r++ {
		cnts := map[byte]int{}
		l, r = filter(r, s1, s2)
		flag := true

		for flag && l <= r {
			switch contain(cnts) {
			case 0:
				return true
			case 1:
				if r < len(s2) && strings.ContainsRune(s1, rune(s2[r])) {
					cnts[s2[r]]++
					r++
				} else if r >= len(s2) {
					return false
				} else {
					flag = false
				}
			case 2:
				cnts[s2[l]]--
				l++
			default:
				continue
			}
		}

	}

	return false
}

func filter(p int, s1, s2 string) (int, int) {
	if !strings.ContainsRune(s1, rune(s2[p])) {
		p++
	}
	return p, p
}

// @lc code=end
