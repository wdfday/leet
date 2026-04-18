/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// Move left pointer to the next alphanumeric character
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Move right pointer to the previous alphanumeric character
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if left < right {
			if toLower(s[left]) != toLower(s[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true

}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// @lc code=end

