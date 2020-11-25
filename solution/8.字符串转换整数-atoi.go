/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

package solution

import (
	"math"
	"strings"
)

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}

	i, sign := 0, 1
	if s[0] == '-' || s[0] == '+' {
		i++
		if s[0] == '-' {
			sign = -1
		}
	}

	ans := 0
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return sign * ans
		}
		if ans > 214748364 ||
			ans == 214748364 && (sign == 1 && s[i] > '7' || sign == -1 && s[i] > '8') {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		ans = 10*ans + int(s[i]-'0')
		i++
	}
	return sign * ans
}

// @lc code=end
