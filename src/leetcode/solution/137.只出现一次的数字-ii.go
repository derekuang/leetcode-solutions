/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

package solution

// @lc code=start
func singleNumber137(nums []int) int {
	seenOnce, seenTwice := 0, 0

	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}

	return seenOnce
}

// @lc code=end
