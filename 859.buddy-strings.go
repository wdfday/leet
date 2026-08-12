/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	a, b := -1, -1
	k := 0

	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			continue
		}

		k++
		if k > 2 {
			return false
		}

		if k == 1 {
			a = i
		} else {
			b = i
		}
	}

	if k == 0 {
		seen := [26]bool{}
		for i := 0; i < len(s); i++ {
			if seen[s[i]-'a'] {
				return true
			}
			seen[s[i]-'a'] = true
		}
		return false
	}

	if k != 2 {
		return false
	}

	return s[a] == goal[b] && s[b] == goal[a]
}

// @lc code=end

