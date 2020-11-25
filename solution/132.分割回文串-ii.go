/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

package solution

// @lc code=start
func minCut132(s string) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	if len(s) <= 1 {
		return 0
	}
	cuts := make([]int, len(s))
	// cuts[0] = 0
	for i := 1; i < len(s); i++ {
		// at most
		cuts[i] = i
		for j := 0; j <= i; j++ {
			if isPalindrome(s[j : i+1]) {
				var challenger int
				if j == 0 {
					challenger = 0
				} else {
					challenger = cuts[j-1] + 1
				}
				cuts[i] = min(cuts[i], challenger)
			}
		}
	}

	return cuts[len(s)-1]
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// @lc code=end
