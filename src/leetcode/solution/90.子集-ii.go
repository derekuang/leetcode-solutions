/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

package solution

import "sort"

// @lc code=start
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	backtracking(nums, []int{}, 0, &ans)
	return
}

func backtracking(nums, set []int, p int, ans *[][]int) {
	temp := make([]int, len(set))
	copy(temp, set)
	*ans = append(*ans, temp)
	for i := p; i < len(nums); i++ {
		if i != p && nums[i] == nums[i-1] {
			continue
		}
		set = append(set, nums[i])
		backtracking(nums, set, i+1, ans)
		set = set[:len(set)-1]
	}
}

// @lc code=end
