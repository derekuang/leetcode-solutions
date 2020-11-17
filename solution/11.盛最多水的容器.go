/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

package solution

// @lc code=start
func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		w, h := r-l, 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		if w*h > ans {
			ans = w * h
		}
	}
	return ans
}

// @lc code=end
