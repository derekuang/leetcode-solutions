/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

package solution

// @lc code=start
func canJump(nums []int) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	flag := 0
	for i := 0; i < len(nums); i++ {
		if i <= flag {
			flag = max(flag, i+nums[i])
			if flag+1 >= len(nums) {
				return true
			}
		}
	}

	return false
}

// @lc code=end
