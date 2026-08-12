/*
 * @lc app=leetcode id=495 lang=golang
 *
 * [495] Teemo Attacking
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		gap := timeSeries[i+1] - timeSeries[i]
		if gap < duration {
			res += gap
		} else {
			res += duration
		}
	}
	return res + duration
}

// @lc code=end

