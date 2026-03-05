/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if minPrice < price {
			profit += price - minPrice
			minPrice = price
		} else {
			minPrice = price
		}
	}
	return profit
}

// @lc code=end

