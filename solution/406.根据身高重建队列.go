/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

package solution

import "math"

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	cnt, checkin := 0, make([]bool, n)
	people = append(people, []int{math.MaxInt64, 0})
	getMin := func() int {
		min := n
		for i, b := range checkin {
			if !b && people[i][0] < people[min][0] {
				min = i
			}
		}
		return min
	}
	for cnt < n-1 {
		p := getMin()
		p0, p1 := people[p][0], people[p][1]
		acc := 0
		for i := 0; i <= p1+acc; i++ {
			if checkin[i] && people[i][0] < p0 {
				acc++
			}
		}
		people[p], people[p1+acc] = people[p1+acc], people[p]
		checkin[p1+acc] = true
		cnt++
	}
	people = people[:n]
	return people
}

// @lc code=end
