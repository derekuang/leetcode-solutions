/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

package solution

// @lc code=start
var m13 = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) (ans int) {
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i : i+1]
		num := m13[c]
		if num >= pre {
			ans += num
		} else {
			ans -= num
		}
		pre = num
	}
	return
}

// @lc code=end
