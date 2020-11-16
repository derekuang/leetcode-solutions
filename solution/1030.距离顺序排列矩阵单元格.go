/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

package solution

// @lc code=start
type coord struct {
	x, y int
}

// BFS
func allCellsDistOrder(R int, C int, r0 int, c0 int) (ans [][]int) {
	q := []coord{}
	q = append(q, coord{r0, c0})
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	visited[r0][c0] = true

	for len(q) > 0 {
		p := q[0]
		x, y := p.x, p.y
		ans = append(ans, []int{x, y})
		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < R &&
				newY >= 0 && newY < C && !visited[newX][newY] {
				q = append(q, coord{newX, newY})
				visited[newX][newY] = true
			}
		}
		q = q[1:]
	}
	return ans
}

// @lc code=end
