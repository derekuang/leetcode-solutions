/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

package solution

// @lc code=start
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	counter := make([]int, 26)
	limit := (n + 1) / 2
	for i := range S {
		if counter[S[i]-'a']+1 > limit {
			return ""
		}
		counter[S[i]-'a']++
	}

	s := make([]byte, n)
	even, odd, half := 0, 1, n/2
	for i, cnt := range counter {
		b := byte(i + 'a')
		for cnt > 0 && cnt <= half && odd < n {
			s[odd] = b
			cnt--
			odd += 2
		}
		for cnt > 0 && even < n {
			s[even] = b
			cnt--
			even += 2
		}
	}
	return string(s)
}

// @lc code=end
