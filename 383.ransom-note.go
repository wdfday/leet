/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, c := range magazine {
		m[c]++
	}

	for _, c := range ransomNote {
		if m[c] == 0 {
			return false
		}
		m[c]--
	}

	return true

}

// @lc code=end

