/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

package solution

// @lc code=start
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	p, q := 0, n-1
	for p < q {
		for p < q && nums[p] != val {
			p++
		}
		for p < q && nums[q] == val {
			q--
		}
		if p < q {
			nums[p], nums[q] = nums[q], nums[p]
		}
	}
	if nums[p] == val {
		return p
	}
	return p + 1
}

// @lc code=end
