/*
 * @lc app=leetcode id=3754 lang=golang
 *
 * [3754] Concatenate Non-Zero Digits and Multiply by Sum I
 */

// @lc code=start
func sumAndMultiply(n int) int64 {
	str := strconv.Itoa(n)
	s, m := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			s = s * 10 + int(str[i] - '0')
			m += int(str[i] - '0')
		}
	}
	return int64(s) * int64(m)
    
}
// @lc code=end

