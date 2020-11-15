/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

package solution

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	stack := make([]int, n)
	for i := 0; i < n; i++ {
		stack[i] = n - 1 - i
	}
	for len(stack) > 0 {
		p := stack[len(stack)-1] // pop
		pHeight, pNum := people[p][0], people[p][1]
		cnt := 0
		for i := 0; i < p; i++ {
			if people[i][0] >= pHeight {
				cnt++
			}
		}
		if cnt == pNum {
			stack = stack[:len(stack)-1]
		} else if cnt < pNum {
			for ; cnt < pNum; p++ {
				if people[p+1][0] >= pHeight {
					cnt++
				}
				people[p], people[p+1] = people[p+1], people[p]
			}
		} else { // cnt > pNum
			stack = stack[:len(stack)-1]
			for ; cnt > pNum; p-- {
				if people[p-1][0] > pHeight {
					cnt--
				} else if people[p-1][0] < pHeight {
					stack = append(stack, p)
				} else {
					cnt--
					stack = append(stack, p)
				}
				people[p-1], people[p] = people[p], people[p-1]
			}
		}
	}
	return people
}

// @lc code=end
