/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

package solution

// @lc code=start
func minPathSum(grid [][]int) int {
	lenSize, colSize := len(grid), len(grid[0])
	dist := make([][]int, lenSize)
	for i := 0; i < len(dist); i++ {
		dist[i] = make([]int, colSize)
	}
	dist[0][0] = grid[0][0]

	for i := 1; i < colSize; i++ {
		dist[0][i] = dist[0][i-1] + grid[0][i]
	}
	for i := 1; i < lenSize; i++ {
		dist[i][0] = dist[i-1][0] + grid[i][0]
	}
	for i := 1; i < lenSize; i++ {
		for j := 1; j < colSize; j++ {
			if dist[i-1][j] < dist[i][j-1] {
				dist[i][j] = dist[i-1][j] + grid[i][j]
			} else {
				dist[i][j] = dist[i][j-1] + grid[i][j]
			}
		}
	}

	return dist[lenSize-1][colSize-1]
}

// @lc code=end
