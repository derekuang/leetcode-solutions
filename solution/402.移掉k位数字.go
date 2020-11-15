/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

package solution

import (
	"strings"
)

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}
	return ans
}

// @lc code=end
