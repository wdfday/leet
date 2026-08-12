/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	n := len(num)

	add := func(a, b string) string {
		i, j := len(a)-1, len(b)-1
		carry := 0
		res := make([]byte, 0, max(len(a), len(b))+1)

		for i >= 0 || j >= 0 || carry > 0 {
			sum := carry

			if i >= 0 {
				sum += int(a[i] - '0')
				i--
			}

			if j >= 0 {
				sum += int(b[j] - '0')
				j--
			}

			res = append(res, byte(sum%10)+'0')
			carry = sum / 10
		}

		// reverse
		for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
			res[l], res[r] = res[r], res[l]
		}

		return string(res)
	}

	check := func(a, b int) bool {
		x := num[:a]
		y := num[a:b]

		// leading zero
		if len(x) > 1 && x[0] == '0' {
			return false
		}
		if len(y) > 1 && y[0] == '0' {
			return false
		}

		i := b

		for i < n {
			z := add(x, y)

			if i+len(z) > n || num[i:i+len(z)] != z {
				return false
			}

			i += len(z)
			x = y
			y = z
		}

		return true
	}

	for a := 1; a <= n-2; a++ {
		for b := a + 1; b <= n-1; b++ {
			if check(a, b) {
				return true
			}
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

