/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

package solution

// @lc code=start
func maxArea(height []int) (ans int) {
	var l, r int
	for r = 1; r < len(height); r++ {
		for l = 0; l < r; l++ {
			v := min11(height[l], height[r]) * (r - l)
			ans = max11(ans, v)
		}
	}
	return ans
}

func min11(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max11(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// @lc code=end
