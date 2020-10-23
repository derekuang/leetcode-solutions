/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

package solution

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	acc := 31
	for num != 0 {
		res += (num & 1) << acc
		num >>= 1
		acc--
	}

	return res
}

// @lc code=end
