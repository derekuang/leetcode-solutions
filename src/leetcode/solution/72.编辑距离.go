/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

package solution

// @lc code=start
// func minDistance(word1 string, word2 string) int {
// 	dp := make([][]int, len(word1)+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, len(word2)+1)
// 	}

// 	for i := 1; i < len(dp); i++ {
// 		dp[i][0] = i
// 	}

// 	for i := 1; i < len(dp[0]); i++ {
// 		dp[0][i] = i
// 	}

// 	for i := 1; i < len(dp); i++ {
// 		for j := 1; j < len(dp[0]); j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else {
// 				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
// 			}
// 		}
// 	}

// 	return dp[len(word1)][len(word2)]
// }

// func min(ret int, nums ...int) int {
// 	for _, num := range nums {
// 		if num < ret {
// 			ret = num
// 		}
// 	}

// 	return ret
// }

// @lc code=end
