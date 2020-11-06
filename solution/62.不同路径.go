/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

package solution

// @lc code=start
func uniquePaths(m int, n int) int {
	paths := make([]int, n)
	for i := 0; i < n; i++ {
		paths[i] = 1
	}

	for i := 0; i < m-1; i++ {
		for j := 1; j < n; j++ {
			paths[j] += paths[j-1]
		}
	}

	return paths[n-1]
}

// @lc code=end
