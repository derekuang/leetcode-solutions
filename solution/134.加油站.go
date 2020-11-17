/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

package solution

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = gas[i] - cost[i]
	}
	acc, p, q := 0, 0, 0
	cnt1, cnt2 := 0, 0
	for acc < n {
		cnt1 += count[q]
		if cnt1 >= 0 {
			q++
			acc++
		} else {
			cnt1 -= count[p]
			cnt2 += count[p]
			p++
			if p > q {
				q = p
				acc++
			} else {
				cnt1 -= count[q]
			}
		}
	}
	if cnt1+cnt2 >= 0 {
		return p
	}
	return -1
}

// @lc code=end
