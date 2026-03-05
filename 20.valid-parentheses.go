/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, byte(c))
		default:
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

