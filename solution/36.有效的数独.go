/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

package solution

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row, square := make([]map[byte]bool, 9), make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = map[byte]bool{}
		square[i] = map[byte]bool{}
	}
	for i := 0; i < 9; i++ {
		column := map[byte]bool{}
		for j := 0; j < 9; j++ {
			cur := board[i][j]
			if cur == '.' {
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
