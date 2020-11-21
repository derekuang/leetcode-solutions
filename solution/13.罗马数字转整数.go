/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

package solution

// @lc code=start
func romanToInt(s string) (ans int) {
	n := len(s)
	for i := 0; i < n; i++ {
		add := 0
		switch s[i] {
		case 'I':
			if i+1 < n && s[i+1] == 'V' {
				add = 4
				i++
			} else if i+1 < n && s[i+1] == 'X' {
				add = 9
				i++
			} else {
				add = 1
			}
		case 'X':
			if i+1 < n && s[i+1] == 'L' {
				add = 40
				i++
			} else if i+1 < n && s[i+1] == 'C' {
				add = 90
				i++
			} else {
				add = 10
			}
		case 'C':
			if i+1 < n && s[i+1] == 'D' {
				add = 400
				i++
			} else if i+1 < n && s[i+1] == 'M' {
				add = 900
				i++
			} else {
				add = 100
			}
		case 'V':
			add = 5
		case 'L':
			add = 50
		case 'D':
			add = 500
		case 'M':
			add = 1000
		}
		ans += add
	}
	return
}

// @lc code=end
