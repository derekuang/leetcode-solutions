/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

package solution

import "strings"

// @lc code=start
func countAndSay(n int) (ans string) {
	if n == 1 {
		return "1"
	}

	ans = "1"
	for i := 2; i <= n; i++ {
		s := strings.Builder{}
		cnt, c := 1, ans[0]
		for i := 1; i <= len(ans); i++ {
			if i == len(ans) || ans[i] != ans[i-1] {
				s.WriteByte(byte(cnt + '0'))
				s.WriteByte(c)
				if i != len(ans) {
					cnt, c = 1, ans[i]
				}
			} else {
				cnt++
			}
		}
		ans = s.String()
	}
	return ans
}

// @lc code=end
