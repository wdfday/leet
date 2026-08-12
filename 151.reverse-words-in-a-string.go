/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	// Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Split the string into words
	words := strings.Fields(s)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words with a single space
	return strings.Join(words, " ")

}

// @lc code=end

