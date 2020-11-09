/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

package solution

// @lc code=start
func kClosest(points [][]int, k int) [][]int {
	distance := func(p int) int {
		return points[p][0]*points[p][0] + points[p][1]*points[p][1]
	}
	qSort := func(l, r int) int {
		pivot := distance(l)
		for l < r {
			dr := distance(r)
			for r > l && dr >= pivot {
				r--
				dr = distance(r)
			}
			points[l], points[r] = points[r], points[l]
			dl := distance(l)
			for l < r && dl <= pivot {
				l++
				dl = distance(l)
			}
			points[l], points[r] = points[r], points[l]
		}
		return l
	}

	pivot := -1
	l, r := 0, len(points)-1
	for pivot != k-1 {
		if pivot < k-1 {
			l = pivot + 1
			pivot = qSort(l, r)
		} else {
			r = pivot
			pivot = qSort(l, r)
		}
	}

	return points[:k]
}

// @lc code=end
