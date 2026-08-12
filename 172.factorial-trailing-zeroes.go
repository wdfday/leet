/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) int {
	res := 0
	ft := 5
	for n >= ft {
		res += n / ft
		ft *= 5
	}
	return res
}

// @lc code=end

