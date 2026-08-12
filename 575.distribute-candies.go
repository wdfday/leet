/*
 * @lc app=leetcode id=575 lang=golang
 *
 * [575] Distribute Candies
 */

// @lc code=start
func distributeCandies(candyType []int) int {
	uniqueCandies := make(map[int]bool)
	for _, candy := range candyType {
		uniqueCandies[candy] = true
	}
	return min(len(uniqueCandies), len(candyType)/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

