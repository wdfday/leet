/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := ""

	// sign
	if (numerator < 0) != (denominator < 0) {
		res += "-"
	}

	// dùng int64 để tránh overflow
	n := int64(numerator)
	d := int64(denominator)

	n = abs64(n)
	d = abs64(d)

	// phần nguyên
	integerPart := n / d
	res += itoa(integerPart)

	remainder := n % d
	if remainder == 0 {
		return res
	}

	res += "."

	// map remainder -> position trong res
	seen := map[int64]int{}

	for remainder != 0 {
		if pos, ok := seen[remainder]; ok {
			// insert '(' tại pos, append ')'
			res = res[:pos] + "(" + res[pos:] + ")"
			return res
		}

		seen[remainder] = len(res)

		remainder *= 10
		res += itoa(remainder / d)
		remainder %= d
	}

	return res
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func itoa(x int64) string {
	return strconv.FormatInt(x, 10)
}

// @lc code=end

