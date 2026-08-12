/*
 * @lc app=leetcode id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 */

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	}
    dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
// @lc code=end

