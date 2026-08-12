/*
 * @lc app=leetcode id=3827 lang=golang
 *
 * [3827] Count Monobit Integers
 */

// @lc code=start
func countMonobit(n int) int {
	res := 1
	k := 1 // k = 2^i - 1: 1, 3, 7, 15, ...
	for k <= n {
		res++
		k = k*2 + 1
	}
	return res
}

// @lc code=end

