/*
 * @lc app=leetcode id=3707 lang=golang
 *
 * [3707] Equal Score Substrings
 */

// @lc code=start
func scoreBalance(s string) bool {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += score(s[i])
	}

	lsum := 0
	for i := 0; i < len(s); i++ {
		lsum += score(s[i])
		if lsum * 2 == sum {
			return true
		} 
	}
	return false
}

func score(b byte) int {
    return int(b-'a') + 1
}
// @lc code=end

