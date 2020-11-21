/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

package solution

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-97]++
	}

	for _, c := range t {
		if cnt[c-97] == 0 {
			return false
		}
		cnt[c-97]--
	}

	return true
}

// @lc code=end
