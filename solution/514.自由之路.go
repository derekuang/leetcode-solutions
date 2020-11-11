/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */

package solution

import "math"

// @lc code=start
func findRotateSteps(ring string, key string) int {
	lRing, lKey := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([][]int, lKey)
	for i := range dp {
		dp[i] = make([]int, lRing)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min514(p, lRing-p)
	}
	for i := 1; i < lKey; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min514(dp[i][j], dp[i-1][k]+min514(abs514(j-k), lRing-abs514(j-k)))
			}
		}
	}
	return min514(dp[lKey-1]...) + lKey
}

func min514(n ...int) int {
	ans := n[0]
	for _, v := range n[1:] {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func abs514(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
