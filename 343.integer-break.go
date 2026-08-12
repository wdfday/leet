/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	res := 1

	for n > 4 {
		res *= 3
		n -= 3
	}

	res *= n
	return res
}

// @lc code=end

