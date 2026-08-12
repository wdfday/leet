/*
 * @lc app=leetcode id=3857 lang=golang
 *
 * [3857] Minimum Cost to Split into Ones
 */

// @lc code=start
func minCost(n int) int {
	if n <= 1 {
		return 0
	}
	if n < 3 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		m := math.MaxInt
		for j := 1; j*j <= i; j++ {
			prod := (i-j)*j + dp[i-j] + dp[j]
			if prod < m {
				m = prod
			}
		}
		dp[i] = m
	}
	return dp[n]
}

// @lc code=end

