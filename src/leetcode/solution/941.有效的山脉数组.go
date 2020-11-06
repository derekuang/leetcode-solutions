/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

package solution

// @lc code=start
const (
	UP      = iota
	DOWN    = iota
	ILLEGAL = iota
)

func validMountainArray(arr []int) bool {
	if len(arr) < 3 || arr[1] < arr[0] {
		return false
	}
	dir := UP
	for i := 1; i < len(arr); i++ {
		if dir == ILLEGAL {
			return false
		}
		if arr[i] == arr[i-1] {
			return false
		} else if dir == UP && arr[i] > arr[i-1] ||
			dir == DOWN && arr[i] < arr[i-1] {
			continue
		} else {
			dir++
		}
	}
	if dir != DOWN {
		return false
	}
	return true
}

// @lc code=end
