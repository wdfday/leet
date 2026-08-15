/*
 * @lc app=leetcode id=3582 lang=golang
 *
 * [3582] Generate Tag for Video Caption
 */

// @lc code=start
func generateTag(caption string) string {
	words := strings.Fields(caption)
	var sb strings.Builder
	sb.WriteByte('#')
	for i, word := range words {
		lower := strings.ToLower(word)
		if i == 0 {
			sb.WriteString(lower)
		} else {
			sb.WriteString(strings.ToUpper(lower[:1]))
			sb.WriteString(lower[1:])
		}
	}

	res := sb.String()
	if len(res) > 100 {
		return res[:100]
	}
	return res
}
// @lc code=end

