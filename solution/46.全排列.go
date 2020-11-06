/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package solution

// @lc code=start
func permute(nums []int) (ans [][]int) {
	backtrack(nums, []int{}, &ans)
	return
}

func backtrack(nums, set []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int(nil), set...))
		return
	}
	size := len(nums)
	var newNums []int
	for i, num := range nums {
		set = append(set, num)
		if size == 1 {
			newNums = []int{}
		} else {
			newNums = append(append([]int(nil), nums[0:i]...), nums[i+1:size]...)
		}
		backtrack(newNums, set, ans)
		set = set[:len(set)-1]
	}
	return
}

// @lc code=end
