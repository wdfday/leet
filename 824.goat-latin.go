/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 */

// @lc code=start
func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	vowels := "aeiouAEIOU"

	res := make([]string, 0, len(words))

	for i, word := range words {
		if !strings.ContainsRune(vowels, rune(word[0])) {
			word = word[1:] + word[:1]
		}
		word += "ma" + strings.Repeat("a", i+1)
		res = append(res, word)
	}

	return strings.Join(res, " ")
}

// @lc code=end

