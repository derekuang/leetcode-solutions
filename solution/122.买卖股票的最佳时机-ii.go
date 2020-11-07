/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

package solution

// @lc code=start
func maxProfit(prices []int) (profits int) {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < len(prices); i++ {
		profits += max(0, prices[i]-prices[i-1])
	}
	return
}

// @lc code=end
