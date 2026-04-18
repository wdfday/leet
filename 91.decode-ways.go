/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[n] = 1 // base case: empty string

	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		// Decode s[i] alone
		if s[i] != '0' {
			dp[i] = dp[i+1]
		}

		// Decode s[i:i+2] together
		num := (s[i]-'0')*10 + (s[i+1] - '0')
		if num >= 10 && num <= 26 {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

// @lc code=end

