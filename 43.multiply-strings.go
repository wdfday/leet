/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m := len(num1)
	n := len(num2)
	result := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')
			sum := product + result[i+j+1]

			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	var sb strings.Builder
	for _, digit := range result {
		if sb.Len() != 0 || digit != 0 {
			sb.WriteByte(byte(digit) + '0')
		}
	}

	return sb.String()

}

// @lc code=end

