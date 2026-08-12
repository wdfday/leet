/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		w := words[i]

		if v, ok := m1[c]; ok {
			if v != w {
				return false
			}
		} else {
			m1[c] = w
		}

		if v, ok := m2[w]; ok {
			if v != c {
				return false
			}
		} else {
			m2[w] = c
		}
	}
	return true
}

// @lc code=end

