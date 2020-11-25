/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

package solution

import "strings"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	prefix := strings.Builder{}
	for i := 0; true; i++ {
		if i >= len(strs[0]) {
			return prefix.String()
		}
		b := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != b {
				return prefix.String()
			}
		}
		prefix.WriteByte(b)
	}
	return ""
}

// @lc code=end
