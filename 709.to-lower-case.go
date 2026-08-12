/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 */

// @lc code=start
func toLowerCase(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] <= 'Z' && b[i] >= 'A' {
			b[i] += 'a' - 'A'
		}
	}

	return string(b)

}

// @lc code=end

