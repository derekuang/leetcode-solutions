/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package solution

// @lc code=start
func subsets(nums []int) (res [][]int) {
	len := len(nums)
	for mask := 0; mask < (1 << len); mask++ {
		set := []int{}
		for i, num := range nums {
			if (mask>>i)&1 > 0 {
				// num should present in the set
				set = append(set, num)
			}
		}
		res = append(res, append([]int(nil), set...))
	}
	return
}

// @lc code=end
