/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	i := 1
	for n >= i {
		n -= i
		i++
	}
	return i - 1
}

// @lc code=end

