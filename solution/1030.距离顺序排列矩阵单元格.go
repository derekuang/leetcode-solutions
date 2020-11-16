/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

package solution

// @lc code=start
// BFS
func allCellsDistOrder(R int, C int, r0 int, c0 int) (ans [][]int) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	q := [][]int{}
	q = append(q, []int{r0, c0})
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	visited[r0][c0] = true

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		ans = append(ans, []int{x, y})
		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < R &&
				newY >= 0 && newY < C && !visited[newX][newY] {
				q = append(q, []int{newX, newY})
				visited[newX][newY] = true
			}
		}
		q = q[1:]
	}
	return ans
}

// @lc code=end
