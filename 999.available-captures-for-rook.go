/*
 * @lc app=leetcode id=999 lang=golang
 *
 * [999] Available Captures for Rook
 */

// @lc code=start
func numRookCaptures(board [][]byte) int {
	res := 0
	x, y := -1, -1
	for i, b := range board {
		for j, v := range b {
			if v == 'R' {
				x, y = i, j
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		i, j := x, y
		for {
			i += d[0]
			j += d[1]
			if i < 0 || i >= 8 || j < 0 || j >= 8 || board[i][j] == 'B' {
				break
			}
			if board[i][j] == 'p' {
				res++
				break
			}
		}
	}

	return res
}

// @lc code=end

