/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

package solution

// @lc code=start
func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	start, end := 0, len(s)-1
	for start < end {
		t := s[start]
		s[start] = s[end]
		s[end] = t
		start++
		end--
	}
}

// @lc code=end
