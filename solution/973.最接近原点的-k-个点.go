/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

package solution

// @lc code=start
func kClosest(points [][]int, K int) (ans [][]int) {
	ans = make([][]int, K)
	for i := 0; i < K; i++ {
		ans[i] = append([]int{}, 10000, 10000)
	}

	pow := func(x int) int {
		return x * x
	}
	for _, p := range points {
		dist := pow(p[0]) + pow(p[1])
		for i, v := range ans {
			iDist := pow(v[0]) + pow(v[1])
			if dist < iDist {
				for j := K - 1; j > i; j-- {
					ans[j] = ans[j-1]
				}
				ans[i] = append([]int{}, p[0], p[1])
				break
			}
		}
	}
	return
}

// @lc code=end
