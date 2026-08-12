/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	words2 := make([]string, len(words))

	for i, word := range words {
		cs := []byte(word)

		l, r := 0, len(cs)-1
		for l < r {
			cs[l], cs[r] = cs[r], cs[l]
			l++
			r--
		}
		words2[i] = string(cs)
	}
	return strings.Join(words2, " ")
}

// @lc code=end

