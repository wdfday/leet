/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	reversed := 0
	for x > 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before multiplying by 10
		if reversed > (math.MaxInt32-digit)/10 {
			return 0 // Overflow, return 0 as per problem statement
		}

		reversed = reversed*10 + digit
	}

	return sign * reversed
}

// @lc code=end

