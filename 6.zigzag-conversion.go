/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	curRow, goingDown := 0, false

	for _, char := range s {
		rows[curRow] += string(char)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result

}

// @lc code=end

