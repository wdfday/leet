/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
func maxProfit(prices []int) int {

	dp := make([][]int, 0)
	for i := 0; i < len(prices); i++ {
		dp = append(dp, make([]int, 5))
	}

	for i := 0; i < len(prices); i++ {
		for j := 0; j < 5; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else if j%2 == 1 {
				if i == 0 {
					dp[i][j] = -prices[i]
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
				}
			} else {
				if i == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
				}
			}
		}
	}

	return max(dp[len(prices)-1][0], max(dp[len(prices)-1][2], dp[len(prices)-1][4]))

}

// @lc code=end

