/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

package solution

// @lc code=start
func reverse(x int) (ans int) {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x != 0 {
		digit := x % 10
		x /= 10
		if ans > 214748364 {
			return 0
		}
		ans = 10*ans + digit
	}
	return sign * ans
}

// @lc code=end
