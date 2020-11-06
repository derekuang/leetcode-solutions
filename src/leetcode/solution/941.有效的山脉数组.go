/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

package solution

// @lc code=start
func validMountainArray(arr []int) bool {
	if len(arr) < 3 || arr[1] < arr[0] || arr[len(arr)-1] > arr[len(arr)-2] {
		return false
	}
	l, r := 0, len(arr)-1
	for l < len(arr) {
		if arr[l+1] <= arr[l] {
			break
		}
		l++
	}
	for r > 0 {
		if arr[r-1] <= arr[r] {
			break
		}
		r--
	}
	return l == r
}

// @lc code=end
