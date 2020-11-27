/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

package solution

import "strings"

// @lc code=start
func isPalindrome125(s string) bool {
	isLegal := func(c byte) bool {
		return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
	}

	s = strings.ToLower(s)
	n := len(s)
	for p, q := 0, n-1; p < q; {
		for ; p < q && !isLegal(s[p]); p++ {
		}
		for ; q > p && !isLegal(s[q]); q-- {
		}
		if p < q && s[p] != s[q] {
			return false
		}
		p++
		q--
	}
	return true
}

// @lc code=end
