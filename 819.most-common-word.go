/*
 * @lc app=leetcode id=819 lang=golang
 *
 * [819] Most Common Word
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	bannedSet := make(map[string]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	words := strings.FieldsFunc(strings.ToLower(paragraph), func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	count := make(map[string]int)
	best, bestCount := "", 0

	for _, w := range words {
		if bannedSet[w] {
			continue
		}
		count[w]++
		if count[w] > bestCount {
			best, bestCount = w, count[w]
		}
	}

	return best
}

// @lc code=end

