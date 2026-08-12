/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	count := make(map[rune]int)

	for _, c := range s {
		count[c]++
	}

	for i, c := range s {
		if count[c] == 1 {
			return i
		}
	}

	return -1

}

// @lc code=end

