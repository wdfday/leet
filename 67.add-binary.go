/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	var carry int
	var result []byte

	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry > 0 {
		var sum int
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		sum += carry
		result = append(result, byte(sum%2)+'0')
		carry = sum / 2
	}

	// Reverse the result
	for i, n := 0, len(result); i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}

	return string(result)
}

// @lc code=end

