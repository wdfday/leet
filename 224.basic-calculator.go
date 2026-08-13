/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
func calculate(s string) int {
	res, num, sign := 0, 0, 1
	stack := []int{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
		case c == '+':
			res += sign * num
			num, sign = 0, 1
		case c == '-':
			res += sign * num
			num, sign = 0, -1
		case c == '(':
			stack = append(stack, res, sign)
			res, sign = 0, 1
		case c == ')':
			res += sign * num
			num = 0
			prevSign := stack[len(stack)-1]
			prevRes := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res = prevRes + prevSign*res
		}
	}
	res += sign * num
	return res
}
// @lc code=end

