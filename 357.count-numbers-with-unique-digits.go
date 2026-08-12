/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	ans := 10      // n = 1
	cur := 9       // số lượng số có đúng 1 chữ số
	available := 9 // còn 9 chữ số để chọn

	for i := 2; i <= n && available > 0; i++ {
		cur *= available
		ans += cur
		available--
	}

	return ans
}

// @lc code=end

