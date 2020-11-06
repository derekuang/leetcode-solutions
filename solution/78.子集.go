/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package solution

// @lc code=start
func subsets(nums []int) (ans [][]int) {
	n, set := len(nums), []int{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		// nums[cur] should be present
		set = append(set, nums[cur])
		dfs(cur + 1)
		// nums[cur] should be absent
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

// @lc code=end
