/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

func climbStairs(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			dp[i] = i
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp[n]

}

// @lc code=end

