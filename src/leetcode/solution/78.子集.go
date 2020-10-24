/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package solution

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	// Get the [i]int{} array
	for i := 1; i <= len(nums); i++ {
		arr, pList := make([]int, i), make([]int, i)

		res = append(res, arr)
	}
}

// @lc code=end
