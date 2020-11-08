/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

package solution

// @lc code=start
type pointList struct {
	index int
	dist  int
	next  *pointList
}

func kClosest(points [][]int, K int) (ans [][]int) {
	pHead := &pointList{-1, 0, nil}

	pow := func(x int) int {
		return x * x
	}
	for i, point := range points {
		d := pow(point[0]) + pow(point[1])
		p, q := pHead, &pointList{i, d, nil}
		acc := 0
		for p.next != nil && acc < K {
			if p.next.dist < d {
				p = p.next
				acc++
			} else {
				break
			}
		}
		q.next = p.next
		p.next = q
	}

	p := pHead.next
	for i := 0; i < K; i++ {
		ans = append(ans, append([]int{}, points[p.index]...))
		p = p.next
	}
	return
}

// @lc code=end
