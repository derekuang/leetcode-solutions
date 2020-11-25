/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

package solution

import "strings"

// @lc code=start
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// @lc code=end
