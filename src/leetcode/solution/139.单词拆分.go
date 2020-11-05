/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

package solution

// @lc code=start
func wordBreak139(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	// "dp[i] = true" means the result would be true while s[:i] as the input of wordBreak
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end
