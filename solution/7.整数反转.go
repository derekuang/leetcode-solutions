/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

package solution

// @lc code=start
func reverse(x int) (ans int) {
	anchor0, anchor1 := 214748364, -214748364
	for x != 0 {
		digit := x % 10
		x /= 10
		if ans > anchor0 || (ans == anchor0 && digit > 7) ||
			ans < anchor1 || (ans == anchor1 && digit > 8) {
			return 0
		}
		ans = 10*ans + digit
	}
	return ans
}

// @lc code=end
