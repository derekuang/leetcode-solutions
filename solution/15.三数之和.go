/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

package solution

import "sort"

// @lc code=start
func threeSum(nums []int) (ans [][]int) {
	cnts := map[int]int{}
	for _, num := range nums {
		cnts[num]++
	}

	unique := []int{}
	for key := range cnts {
		unique = append(unique, key)
	}
	sort.Ints(unique)

	for i := 0; i < len(unique); i++ {
		if unique[i] == 0 && cnts[unique[i]] >= 3 {
			ans = append(ans, []int{0, 0, 0})
			continue
		}
		for j := i + 1; j < len(unique); j++ {
			if 2*unique[i]+unique[j] == 0 && cnts[unique[i]] >= 2 {
				ans = append(ans, []int{unique[i], unique[i], unique[j]})
			} else if unique[i]+2*unique[j] == 0 && cnts[unique[j]] >= 2 {
				ans = append(ans, []int{unique[i], unique[j], unique[j]})
			} else {
				rest := -unique[i] - unique[j]
				if rest > unique[j] && cnts[rest] >= 1 {
					ans = append(ans, []int{unique[i], unique[j], rest})
				}
			}
		}
	}
	return
}

// @lc code=end
