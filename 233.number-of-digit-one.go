/*
 * @lc app=leetcode id=233 lang=golang
 *
 * [233] Number of Digit One
 */

// @lc code=start
func countDigitOne(n int) int {
	count := 0
	factor := 1

	for factor <= n {
		lower := n % factor
		curr := (n / factor) % 10
		higher := n / (factor * 10)

		switch curr {
		case 0:
			count += higher * factor
		case 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		factor *= 10
	}

	return count
}

// @lc code=end

