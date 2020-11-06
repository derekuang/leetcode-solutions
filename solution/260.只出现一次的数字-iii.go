/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

package solution

// @lc code=start
func singleNumber260(nums []int) []int {
	bitmask := 0
	for _, num := range nums {
		bitmask ^= num
	}

	// find the rightmost 1-bit
	diff := bitmask & -bitmask

	res := 0
	for _, num := range nums {
		if (diff & num) != 0 {
			res ^= num
		}
	}

	return []int{res, res ^ bitmask}
}

// @lc code=end
