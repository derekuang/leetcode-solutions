/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

package solution

// @lc code=start
func fib(N int) int {
	if N <= 1 {
		return N
	}
	N1, N2 := 0, 1
	var res int
	for i := 2; i <= N; i++ {
		res = N1 + N2
		N1 = N2
		N2 = res
	}

	return res
}

// @lc code=end
