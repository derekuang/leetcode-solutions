/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package solution

// @lc code=start
func subsets(nums []int) [][]int {
	len := len(nums)
	res := [][]int{}
	res = append(res, []int{})
	// Get the [i]int{} array
	for i := 1; i <= len; i++ {
		// Pointer list save every subscript of nums that make up the new set
		pList := make([]int, i)
		for j := 0; j < i; j++ {
			pList[j] = j
		}
		// The last sub set of [i]int{} must be nums[len-i : len]
		for pList[0] <= len-i {
			newSet := make([]int, i)
			// Assign value to the new set
			for j, index := range pList {
				newSet[j] = nums[index]
			}
			res = append(res, newSet)
			// Get the subscripts of next set with backtrack
			p := i - 1
			for p >= 0 && pList[p] == len+p-i {
				p--
			}
			if p >= 0 {
				pList[p]++
				for p++; p < i; p++ {
					pList[p] = pList[p-1] + 1
				}
			} else {
				pList[0] = len
			}
		}
	}
	return res
}

// @lc code=end
