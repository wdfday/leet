/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if c, ok := m[s[i]]; ok && c >= left {
			left = c + 1
		}
		m[s[i]] = i
		res = max(i-left+1, res)
	}
	return res
}

// @lc code=end