/*
 * @lc app=leetcode id=1025 lang=golang
 *
 * [1025] Divisor Game
 */

// @lc code=start
func divisorGame(n int) bool {

	dp := make([]bool, n + 1)

	for i := 2; i <= n; i++ {
		for k := 1; i % k == 0 && k < i; k++ {
			if dp[i-k] == false {
				dp[i] = true
			}
		}
	}
	
    return dp[n]
}
// @lc code=end

func divisorGame(n int) bool {
    return n%2 == 0
}

