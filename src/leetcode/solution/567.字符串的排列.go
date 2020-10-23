/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

package solution

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}

	matches, cnts, needs := 0, map[byte]int{}, map[byte]int{}
	for i := 0; i < s1Len; i++ {
		needs[s1[i]]++
		cnts[s2[i]]++
	}
	for i := byte('a'); i <= 'z'; i++ {
		if needs[i] == cnts[i] {
			matches++
		}
	}

	l, r := 0, s1Len-1
	for r < s2Len {
		if matches == 26 {
			return true
		}
		if r == s2Len-1 {
			return false
		}
		cnts[s2[l]]--
		l++
		if cnts[s2[l-1]] == needs[s2[l-1]] {
			matches++
		} else if cnts[s2[l-1]] == needs[s2[l-1]]-1 {
			matches--
		}
		cnts[s2[r+1]]++
		r++
		if cnts[s2[r]] == needs[s2[r]] {
			matches++
		} else if cnts[s2[r]] == needs[s2[r]]+1 {
			matches--
		}
	}

	return false
}

// @lc code=end
