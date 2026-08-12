/*
 * @lc app=leetcode id=3813 lang=golang
 *
 * [3813] Vowel-Consonant Score
 */

// @lc code=start
func vowelConsonantScore(s string) int {
	v, c := 0, 0
	for _, ch := range s {
		if strings.ContainsRune("aeiou", ch) {
			v++
		} else if ch >= 'a' && ch <= 'z' {
			c++
		}
	}
	if c == 0 {
		return 0
	}
	return v / c
}

// @lc code=end

