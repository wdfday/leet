/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {

	if divisor == 0 {
		return math.MaxInt32
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend < 0) != (divisor < 0)

	dividend, divisor = abs(dividend), abs(divisor)
	result := 0

	for dividend >= divisor {
		temp, multiple := divisor, 1
		for dividend >= temp {
			dividend -= temp
			result += multiple
			temp <<= 1
			multiple <<= 1
		}
	}

	if negative {
		result = -result
	}

	return result

}

// @lc code=end

