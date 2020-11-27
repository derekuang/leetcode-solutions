/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

package solution

// @lc code=start
func fourSumCount(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}

// @lc code=end
