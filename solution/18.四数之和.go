/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

package solution

import "sort"

// @lc code=start
func fourSum(nums []int, target int) (ans [][]int) {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	unique := []int{}
	for key := range counter {
		unique = append(unique, key)
	}
	n := len(unique)
	sort.Ints(unique)

	for i := 0; i < n; i++ {
		n1 := unique[i]
		if n1*4 == target && counter[n1] >= 4 {
			ans = append(ans, []int{n1, n1, n1, n1})
			continue
		}
		for j := i + 1; j < n; j++ {
			n2 := unique[j]
			if n1*3+n2 == target && counter[n1] >= 3 {
				ans = append(ans, []int{n1, n1, n1, n2})
			} else if n1*2+n2*2 == target && counter[n1] >= 2 && counter[n2] >= 2 {
				ans = append(ans, []int{n1, n1, n2, n2})
			} else if n1+n2*3 == target && counter[n2] >= 3 {
				ans = append(ans, []int{n1, n2, n2, n2})
			}
			for k := j + 1; k < n; k++ {
				n3 := unique[k]
				if n1*2+n2+n3 == target && counter[n1] >= 2 {
					ans = append(ans, []int{n1, n1, n2, n3})
				} else if n1+n2*2+n3 == target && counter[n2] >= 2 {
					ans = append(ans, []int{n1, n2, n2, n3})
				} else if n1+n2+n3*2 == target && counter[n3] >= 2 {
					ans = append(ans, []int{n1, n2, n3, n3})
				}
				n4 := target - n1 - n2 - n3
				if n4 > n3 && counter[n4] >= 1 {
					ans = append(ans, []int{n1, n2, n3, n4})
				}
			}
		}
	}
	return ans
}

// @lc code=end
