/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	cs := make([]byte, 0)
	ct := make([]byte, 0)

	for _, v := range s {
		if v == '#' {
			if len(cs) > 0 {
				cs = cs[:len(cs)-1]
			}
		} else {
			cs = append(cs, byte(v))
		}
	}
	for _, v := range t {
		if v == '#' {
			if len(ct) > 0 {
				ct = ct[:len(ct)-1]
			}
		} else {
			ct = append(ct, byte(v))
		}
	}

	return string(cs) == string(ct)
}

// @lc code=end

