/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var bfs func(x, y int) 
	bfs = func(x, y int) {
		board[x][y] = '#'

		for _, dir := range dirs {
			u, v := x + dir[0], y + dir[1]
			if u < 0 || v < 0 || u >= m || v >= n {
				continue
			}
			if board[u][v] != 'O' {
				continue
			}
			bfs(u, v)
		}
	}


	for i := range m {
		for j := range n {
			if i != 0 && i != m-1 && j != 0 && j != n-1 {
				continue
			}

			if board[i][j] == 'O' {
				bfs(i, j)
			}
		}
	}

	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := range m {
		for j := range n {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}


}

// @lc code=end

