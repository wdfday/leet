/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */

// @lc code=start
func judgeCircle(moves string) bool {

	p := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			p += 1000000
		case 'D':
			p -= 1000000
		case 'L':
			p += 1
		case 'R':
			p -= 1
		}
	}
	return p == 0
}

// @lc code=end

