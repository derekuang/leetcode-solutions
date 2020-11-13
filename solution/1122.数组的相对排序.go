/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

package solution

import (
	"sort"
)

// @lc code=start
var m map[int]int

type slice1122 []int

func (arr slice1122) Len() int {
	return len(arr)
}
func (arr slice1122) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr slice1122) Less(i, j int) bool {
	if m[arr[i]] == m[arr[j]] {
		return arr[i] < arr[j]
	}
	return m[arr[i]] > m[arr[j]]
}

func relativeSortArray(arr1, arr2 []int) []int {
	m = map[int]int{}
	n := len(arr2)
	for i, v := range arr2 {
		m[v] = n - i
	}
	sort.Sort(slice1122(arr1))
	return arr1
}

// @lc code=end
