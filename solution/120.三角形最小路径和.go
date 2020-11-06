/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

package solution

// @lc code=start
func minimumTotal(triangle [][]int) int {
	dist := make([][]int, len(triangle))
	for i, t := range triangle {
		dist[i] = make([]int, len(t))
	}

	for i, layer := range triangle {
		for j, spot := range layer {
			shortestPath := getShortestPath(i, j, dist)
			dist[i][j] = shortestPath + spot
		}
	}

	res := dist[len(dist)-1][0]
	for _, distance := range dist[len(dist)-1] {
		if distance < res {
			res = distance
		}
	}

	return res
}

func getShortestPath(x, y int, dist [][]int) int {
	if x == 0 && y == 0 {
		return 0
	}

	switch y {
	case 0:
		return dist[x-1][y]
	case x:
		return dist[x-1][y-1]
	default:
		if dist[x-1][y] < dist[x-1][y-1] {
			return dist[x-1][y]
		}
		return dist[x-1][y-1]
	}
}

// @lc code=end
