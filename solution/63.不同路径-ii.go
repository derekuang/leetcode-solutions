/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

package solution

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height, width := len(obstacleGrid), len(obstacleGrid[0])
	paths := make([]int, width)
	if obstacleGrid[0][0] == 0 {
		paths[0] = 1
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if obstacleGrid[i][j] == 1 {
				paths[j] = 0
			} else if j != 0 {
				paths[j] += paths[j-1]
			}
		}
	}

	return paths[width-1]
}

// @lc code=end
