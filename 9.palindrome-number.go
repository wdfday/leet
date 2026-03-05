/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before multiplying by 10
		if reversed > (math.MaxInt32-digit)/10 {
			return false // Overflow, not a palindrome
		}

		reversed = reversed*10 + digit
	}

	return original == reversed
}

// @lc code=end

