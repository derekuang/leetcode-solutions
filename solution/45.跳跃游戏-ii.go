/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

package solution

// @lc code=start
func jump(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	step, pos, end, length := 0, 0, 0, len(nums)
	for i := 0; i < length-1; i++ {
		pos = max(pos, i+nums[i])
		if pos >= length-1 {
			return step + 1
		}
		if i == end {
			end = pos
			step++
		}
	}

	return step
}

// @lc code=end
