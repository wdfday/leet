/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

// @lc code=start
func findWords(words []string) []string {
	rowOf := make(map[rune]int)
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

	for i, row := range rows {
		for _, c := range row {
			rowOf[c] = i
		}
	}

	var res []string

	for _, word := range words {
		lower := strings.ToLower(word)
		targetRow := rowOf[rune(lower[0])]
		valid := true

		for _, c := range lower {
			if rowOf[c] != targetRow {
				valid = false
				break
			}
		}

		if valid {
			res = append(res, word)
		}
	}

	return res
}

// @lc code=end

