/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return true
	}

	ss := s + s
	if strings.Contains(ss[1:len(ss)-1], goal) {
		return true
	}

	return false
}

// @lc code=end

