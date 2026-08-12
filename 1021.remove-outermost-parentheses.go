/*
 * @lc app=leetcode id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 */

// @lc code=start
func removeOuterParentheses(s string) string {	
	bytes := make([]byte, 0)
	m := 0
	for i := 0; i < len(s); i++ {
		if s[i] =='(' {
			m++
			if m != 1 {
				bytes = append(bytes, '(')
			}
		} else {
			m--
			if m != 0 {
				bytes = append(bytes, ')')
			}
		}
	}
    return string(bytes)
}
// @lc code=end

