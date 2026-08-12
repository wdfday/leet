/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var line []string
	lineLen := 0

	for i := 0; i < len(words); i++ {
		w := words[i]

		// nếu thêm word này vào vượt width (tính cả spaces giữa words)
		if lineLen+len(w)+len(line) > maxWidth {
			res = append(res, buildLine(line, lineLen, maxWidth, false))
			line = []string{}
			lineLen = 0
		}

		line = append(line, w)
		lineLen += len(w)
	}

	// last line
	res = append(res, buildLine(line, lineLen, maxWidth, true))

	return res
}

func buildLine(words []string, wordsLen, maxWidth int, last bool) string {
	var sb strings.Builder

	// case last line hoặc 1 word → left justify
	if last || len(words) == 1 {
		for i, w := range words {
			sb.WriteString(w)
			if i != len(words)-1 {
				sb.WriteString(" ")
			}
		}
		// pad spaces
		sb.WriteString(strings.Repeat(" ", maxWidth-sb.Len()))
		return sb.String()
	}

	// normal case → distribute spaces
	gaps := len(words) - 1
	totalSpaces := maxWidth - wordsLen

	base := totalSpaces / gaps
	extra := totalSpaces % gaps

	for i, w := range words {
		sb.WriteString(w)
		if i < gaps {
			spaces := base
			if i < extra {
				spaces++
			}
			sb.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return sb.String()
}

// @lc code=end

