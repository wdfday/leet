/*
 * @lc app=leetcode id=917 lang=golang
 *
 * [917] Reverse Only Letters
 */

// @lc code=start
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	lo, hi := 0, len(b)-1
	isLetter := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}
	for lo < hi {
		if !isLetter(b[lo]) {
			lo++
		} else if !isLetter(b[hi]) {
			hi--
		} else {
			b[lo], b[hi] = b[hi], b[lo]
			lo++
			hi--
		}
	}
	return string(b)
}

// @lc code=end

