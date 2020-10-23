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

	cnts, needs := map[byte]int{}, map[byte]int{}
	for i := 0; i < s1Len; i++ {
		needs[s1[i]]++
		cnts[s2[i]]++
	}

	isContain := func() bool {
		for k, v := range needs {
			if cnts[k] != v {
				return false
			}
		}
		return true
	}

	l, r := 0, s1Len
	for r <= s2Len {
		if isContain() {
			return true
		} else if r < s2Len {
			cnts[s2[l]]--
			cnts[s2[r]]++
			l++
			r++
		} else {
			return false
		}
	}

	return false
}

// @lc code=end
