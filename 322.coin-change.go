/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	for i := range amount + 1 {
		if i != 0 {
			dp[i] = 1e9
		}
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], 1 + dp[i - coin])
		}
	}

	if dp[amount] == 1e9 {
		return -1
	} 
	return dp[amount]
}
// @lc code=end

