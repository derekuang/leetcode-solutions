/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

package solution

// @lc code=start
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n <= 1 || k == 0 {
		return
	}

	for i, cnt := 0, 0; cnt < n; i++ {
		pre, cur := nums[i], (i+k)%n
		for cur != i {
			pre, nums[cur] = nums[cur], pre
			cur = (cur + k) % n
			cnt++
		}
		nums[cur] = pre
		cnt++
	}
}

// @lc code=end
