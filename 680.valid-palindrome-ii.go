/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */

// @lc code=start
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			// If a mismatch is found, check both branches:
			// 1. Skip the left character: check s[l+1...r]
			// 2. Skip the right character: check s[l...r-1]
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
	}

	return true
}

// Helper function to check if a substring is a strict palindrome
func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end

