/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

package solution

// @lc code=start
func maxProfit(prices []int) (profits int) {
	profit := 0
	for buy, sell := 0, 1; sell < len(prices); sell++ {
		if prices[sell] <= prices[sell-1] {
			buy = sell
			profits += profit
			profit = 0
		} else if prices[sell]-prices[buy] > profit {
			profit = prices[sell] - prices[buy]
		}
	}
	profits += profit
	return
}

// @lc code=end
