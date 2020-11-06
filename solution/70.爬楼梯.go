/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

package solution

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	oneStep, twoStep := 1, 2
	res := 0
	for i := 3; i <= n; i++ {
		res = oneStep + twoStep
		oneStep = twoStep
		twoStep = res
	}

	return res
}

// @lc code=end
