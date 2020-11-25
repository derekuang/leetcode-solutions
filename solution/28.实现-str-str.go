/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

package solution

// @lc code=start
// KMP
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if haystack == "" || len(haystack) < len(needle) {
		return -1
	}

	next := getNext(needle)
	i, j := 0, 0
	n1, n2 := len(haystack), len(needle)
	for i < n1 && j < n2 {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == n2 {
		return i - n2
	}
	return -1
}

func getNext(s string) []int {
	n := len(s)
	next := make([]int, n)
	next[0] = -1
	for p, q := 0, 1; q < n; {
		if s[p] == s[q] {
			next[q] = next[p]
		} else {
			next[q] = p
			for p != -1 && s[p] != s[q] {
				p = next[p]
			}
		}
		q++
		p++
	}
	return next
}

// @lc code=end
