/*
 * @lc app=leetcode id=3856 lang=golang
 *
 * [3856] Trim Trailing Vowels
 */

// @lc code=start
func trimTrailingVowels(s string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i := len(s) - 1; i >= 0; i-- {
		if !vowels[s[i]] {
			return s[:i+1]
		}
	}
	return ""
}

// @lc code=end

