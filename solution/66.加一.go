/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

package solution

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	sum := 1 + digits[n-1]
	carry := sum / 10
	digits[n-1] = sum % 10
	for i := n - 2; i >= 0; i-- {
		if carry == 0 {
			break
		}
		sum = 1 + digits[i]
		digits[i], carry = sum%10, sum/10
	}
	if carry == 0 {
		return digits
	}
	return append([]int{1}, digits...)
}

// @lc code=end
