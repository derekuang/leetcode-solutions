/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

package solution

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	// First flip upside down
	var p, q int
	if n%2 == 1 {
		p, q = n/2-1, n/2+1
	} else {
		p, q = n/2-1, n/2
	}
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[p][j], matrix[q][j] = matrix[q][j], matrix[p][j]
		}
		p--
		q++
	}
	// Then flip along the positive diagonal
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end
