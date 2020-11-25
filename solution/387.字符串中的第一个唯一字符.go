/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

package solution

// @lc code=start
func firstUniqChar(s string) int {
	counter := [26]int{}
	for _, c := range s {
		counter[c-'a']++
	}
	for i, c := range s {
		if counter[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
