/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

package solution

// @lc code=start
func titleToNumber(s string) (acc int) {
	for _, b := range s {
		acc = acc*26 + int(byte(b)-64)
	}
	return
}

// @lc code=end
