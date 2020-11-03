/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

package solution

import "sort"

// @lc code=start
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	backtracking47(nums, []int{}, &ans)
	return
}

func backtracking47(nums, set []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, set...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		set = append(set, nums[i])
		var newNums []int
		if len(nums) == 1 {
			newNums = []int{}
		} else {
			newNums = append(append([]int{}, nums[0:i]...), nums[i+1:]...)
		}
		backtracking47(newNums, set, ans)
		set = set[:len(set)-1]
	}
}

// @lc code=end
