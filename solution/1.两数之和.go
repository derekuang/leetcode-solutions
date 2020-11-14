/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

package solution

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if m[target-v] != 0 {
			return []int{m[target-v] - 1, i}
		}
		m[v] = i + 1
	}
	return []int{}
}

// @lc code=end
