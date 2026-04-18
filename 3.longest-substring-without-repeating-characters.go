/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if index, ok := charIndex[s[right]]; ok && index >= left {
			left = index + 1
		}
		charIndex[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// @lc code=end