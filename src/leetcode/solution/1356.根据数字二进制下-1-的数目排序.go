/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

package solution

import "sort"

// @lc code=start
type arr1356 []int

var hash map[int]int

func (arr arr1356) Len() int {
	return len(arr)
}

func (arr arr1356) Less(i, j int) bool {
	if hash[arr[i]] == hash[arr[j]] {
		return arr[i] < arr[j]
	}
	return hash[arr[i]] < hash[arr[j]]
}

func (arr arr1356) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sortByBits(arr []int) []int {
	hash = make(map[int]int, len(arr))
	for _, v := range arr {
		t, cnt := v, 0
		for t != 0 {
			if t&1 == 1 {
				cnt++
			}
			t >>= 1
		}
		hash[v] = cnt
	}

	sort.Stable(arr1356(arr))

	return arr
}

// @lc code=end
