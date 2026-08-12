/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		cs, ct := s[i], t[i]
		if v, ok := m1[cs]; ok {
			if v != ct {
				return false
			}
		} else {
			m1[cs] = ct
		}
		if v, ok := m2[ct]; ok {
			if v != cs {
				return false
			}
		} else {
			m2[ct] = cs
		}
	}
	return true
}

// @lc code=end

