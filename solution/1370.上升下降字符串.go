/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

package solution

// @lc code=start
func sortString(s string) string {
	sLen := len(s)
	counter := ['z' + 1]int{}
	for _, c := range s {
		counter[c]++
	}

	ans := []byte{}
	for len(ans) < sLen {
		for c := 'a'; c <= 'z'; c++ {
			if counter[c] > 0 {
				counter[c]--
				ans = append(ans, byte(c))
			}
		}
		for c := 'z'; c >= 'a'; c-- {
			if counter[c] > 0 {
				counter[c]--
				ans = append(ans, byte(c))
			}
		}
	}
	return string(ans)
}

// @lc code=end
