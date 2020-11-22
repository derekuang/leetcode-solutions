/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

package solution

// @lc code=start
func wordBreak(s string, wordDict []string) (ans []string) {
	isContainWord := func(w string) bool {
		for _, word := range wordDict {
			if w == word {
				return true
			}
		}
		return false
	}

	if s == "" {
		return
	}

	n := len(s)
	for i := 1; i <= n; i++ {
		w1 := s[:i]
		if isContainWord(w1) {
			rest := wordBreak(s[i:], wordDict)
			if len(rest) > 0 {
				for _, w := range rest {
					w1 = w1 + " " + w
					ans = append(ans, w1)
				}
			} else if i == n {
				ans = append(ans, w1)
			}
		}
	}
	return
}

// @lc code=end
