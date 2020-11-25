/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

package solution

// @lc code=start
func isPalindrome(s string) bool {
	isLegal := func(c byte) bool {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') {
			return true
		}
		return false
	}
	abs := func(x, y byte) byte {
		if x > y {
			return x - y
		}
		return y - x
	}
	isEqual := func(b1, b2 byte) bool {
		if b1 < 'A' && b2 < 'A' && b1 == b2 {
			return true
		}
		if b1 >= 'A' && b2 >= 'A' && (b1 == b2 || abs(b1, b2) == 32) {
			return true
		}
		return false
	}

	n := len(s)
	for p, q := 0, n-1; p < q; {
		for ; p < q && !isLegal(s[p]); p++ {
		}
		for ; q > p && !isLegal(s[q]); q-- {
		}
		if p < q && !isEqual(s[p], s[q]) {
			return false
		}
		p++
		q--
	}
	return true
}

// @lc code=end
