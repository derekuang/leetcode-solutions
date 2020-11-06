/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

package solution

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	height, width := len(matrix), len(matrix[0])
	head, tail := 0, height*width-1
	for head <= tail {
		mid := (head + tail) / 2
		val := matrix[mid/width][mid%width]
		if target == val {
			return true
		} else if target > val {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}

	return false
}

// @lc code=end
