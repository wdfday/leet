/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	res := make([]string, 0)
	m := make(map[string]int)

	ws1 := strings.Split(s1, " ")
	ws2 := strings.Split(s2, " ")
	for _, w := range ws1 {
		m[w]++
	}
	for _, w := range ws2 {
		m[w]++
	}

	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// @lc code=end

