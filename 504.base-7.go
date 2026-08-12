/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	negative := false
	if num < 0 {
		negative = true
		num = -num
	}

	ans := make([]byte, 0)
	for num > 0 {
		ans = append(ans, byte(num%7)+'0')
		num /= 7
	}

	if negative {
		ans = append(ans, '-')
	}

	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}

	return string(ans)
}

// @lc code=end

