/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

package solution

// @lc code=start
func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	res := []int{}
	if sLen < pLen {
		return res
	}

	cnts, needs := map[byte]int{}, map[byte]int{}
	isContain := func() bool {
		for k, v := range needs {
			if cnts[k] != v {
				return false
			}
		}
		return true
	}

	for i := 0; i < pLen; i++ {
		needs[p[i]]++
		cnts[s[i]]++
	}
	if isContain() {
		res = append(res, 0)
	}

	for l, r := 0, pLen; r < sLen; l, r = l+1, r+1 {
		cnts[s[l]]--
		cnts[s[r]]++
		if isContain() {
			res = append(res, l+1)
		}
	}

	return res
}

// @lc code=end
