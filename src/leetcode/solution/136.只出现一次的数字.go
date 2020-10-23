/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

package solution

// @lc code=start

func singleNumber136(nums []int) int {
	e := 0
	for _, num := range nums {
		e ^= num
	}
	return e
}

// @lc code=end
