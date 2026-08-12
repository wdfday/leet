/*
 * @lc app=leetcode id=640 lang=golang
 *
 * [640] Solve the Equation
 */

// @lc code=start
func solveEquation(equation string) string {
	parse := func(s string) (int, int) {
		coefX, constVal := 0, 0
		sign, num := 1, 0
		hasNum := false

		for i := 0; i <= len(s); i++ {
			var c byte
			if i < len(s) {
				c = s[i]
			} else {
				c = '#'
			}

			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
				hasNum = true
			} else if c == 'x' {
				if !hasNum {
					num = 1
				}
				coefX += sign * num
				num = 0
				hasNum = false
			} else {
				constVal += sign * num
				num = 0
				hasNum = false

				if c == '+' {
					sign = 1
				} else if c == '-' {
					sign = -1
				}
			}
		}
		return coefX, constVal
	}

	leftX, leftC := parse(splitLeft(equation))
	rightX, rightC := parse(splitRight(equation))

	A := leftX - rightX
	B := rightC - leftC

	if A == 0 {
		if B == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return fmt.Sprintf("x=%d", B/A)
}

func splitLeft(equation string) string {
	for i := 0; i < len(equation); i++ {
		if equation[i] == '=' {
			return equation[:i]
		}
	}
	return ""
}

func splitRight(equation string) string {
	for i := 0; i < len(equation); i++ {
		if equation[i] == '=' {
			return equation[i+1:]
		}
	}
	return ""
}

// @lc code=end

