/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

package solution

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row, square := make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		square[i], row[i] = make([]bool, 10), make([]bool, 10)
	}
	for i := 0; i < 9; i++ {
		column := make([]bool, 10)
		for j := 0; j < 9; j++ {
			cur := board[i][j] - 48
			if cur+48 == '.' {
				continue
			} else if column[cur] || row[j][cur] || square[i/3*3+j/3][cur] {
				return false
			} else {
				column[cur] = true
				row[j][cur] = true
				square[i/3*3+j/3][cur] = true
			}
		}
	}
	return true
}

// @lc code=end
