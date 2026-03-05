/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	maxLen := 0
	left, right := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			currLen := 2 * right
			if currLen > maxLen {
				maxLen = currLen
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}

		if left == right {
			currLen := 2 * left
			if currLen > maxLen {
				maxLen = currLen
			}
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}

// @lc code=end

