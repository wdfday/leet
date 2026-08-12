/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

// @lc code=start
func checkRecord(s string) bool {
	totalAbsent := 0
	consecLate := 0
	for _, c := range s {
		if c == 'A' {
			totalAbsent++
			if totalAbsent >= 2 {
				return false
			}
		}
		if c == 'L' {
			consecLate++
			if consecLate > 2 {
				return false
			}
		} else {
			consecLate = 0
		}
	}

	return true

}

// @lc code=end

