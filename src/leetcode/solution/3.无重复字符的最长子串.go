/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

package solution

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	res, cnts := 0, map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnts[s[r]]++
		if cnts[s[r]] == 1 {
			if r-l+1 > res {
				res++
			}
		} else {
			for cnts[s[r]] > 1 {
				cnts[s[l]]--
				l++
			}
		}
	}

	return res
}

// @lc code=end
