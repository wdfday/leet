/*
 * @lc app=leetcode id=3560 lang=golang
 *
 * [3560] Find Minimum Log Transportation Cost
 */

// @lc code=start
func minCuttingCost(n int, m int, k int) int64 {
	if m <= k && n <= k {
		return 0 
	} else if m <= k {
		// if n > 2 * k {
		// 	return int64((n - 2 * k) * k)
		// } else {
			return int64((n - k) * k) 
		// }
	} else if n <= k {
		// if m > 2 * k {
		// 	return int64((m - 2 * k) * k * k)
		// } else {
			return int64((m - k) * k )
		// }
	}
    return 0
}
// @lc code=end

