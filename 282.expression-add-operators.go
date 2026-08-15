/*
 * @lc app=leetcode id=282 lang=golang
 *
 * [282] Expression Add Operators
 */

// @lc code=start
func addOperators(num string, target int) []string {
	res := []string{}
	n := len(num)

	var backtrack func(idx int, expr string, curVal, prevOperand int64)
	backtrack = func(idx int, expr string, curVal, prevOperand int64) {
		if idx == n {
			if curVal == int64(target) {
				res = append(res, expr)
			}
			return
		}

		for i := idx; i < n; i++ {
			if i > idx && num[idx] == '0' {
				break
			}

			curStr := num[idx : i+1]
			curNum, _ := strconv.ParseInt(curStr, 10, 64)

			if idx == 0 {
				backtrack(i+1, curStr, curNum, curNum)
			} else {
				backtrack(i+1, expr+"+"+curStr, curVal+curNum, curNum)
				backtrack(i+1, expr+"-"+curStr, curVal-curNum, -curNum)
				backtrack(i+1, expr+"*"+curStr,
					curVal-prevOperand+prevOperand*curNum,
					prevOperand*curNum)
			}
		}
	}

	backtrack(0, "", 0, 0)
	return res
}
// @lc code=end

func addOperators(num string, target int) []string {
	n := len(num)
	res := []string{}
	buf := make([]byte, n*2)
	var dfs func(idx int, val, last int64, pos int)
	dfs = func(idx int, val, last int64, pos int) {
		if idx == n {
			if val == int64(target) {
				res = append(res, string(buf[:pos]))
			}
			return
		}
		var curr int64 = 0
		for i := idx; i < n; i++ {
			if i > idx && num[idx] == '0' {
				break
			}
			curr = curr*10 + int64(num[i]-'0')
			digits := i - idx + 1

			if idx == 0 {
				for j := 0; j < digits; j++ {
					buf[j] = num[idx+j]
				}
				dfs(i+1, curr, curr, digits)
			} else {
				p := pos
				buf[p] = '+'
				for j := 0; j < digits; j++ {
					buf[p+1+j] = num[idx+j]
				}
				dfs(i+1, val+curr, curr, p+1+digits)

				buf[p] = '-'
				dfs(i+1, val-curr, -curr, p+1+digits)

				buf[p] = '*'
				dfs(i+1, val-last+last*curr, last*curr, p+1+digits)
			}
		}
	}
	dfs(0, 0, 0, 0)
	return res
}