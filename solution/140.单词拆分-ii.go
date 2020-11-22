/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

package solution

import "strings"

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	isContainWord := func(w string) bool {
		for _, word := range wordDict {
			if w == word {
				return true
			}
		}
		return false
	}

	n := len(s)
	if n == 0 {
		return []string{}
	}

	prev := map[int][]string{}
	prev[0] = []string{""}

	for i := 1; i <= n; i++ {
		for pre := range prev {
			cur := s[pre:i]
			if isContainWord(cur) {
				for _, word := range prev[pre] {
					prev[i] = append(prev[i], word+" "+cur)
				}
			}
		}
	}

	ans := prev[n]
	for i := range ans {
		ans[i] = strings.TrimLeft(ans[i], " ")
	}
	return prev[n]
}

// @lc code=end
