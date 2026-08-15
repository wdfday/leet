/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLife(board [][]int)  {
	m, n := len(board), len(board[0])

	dirs := [][2]int{
		{-1,-1}, {-1, 0}, {-1, 1},
		{0,-1}, {0, 1},
		{1,-1}, {1, 0}, {1, 1},

	}
	life := func(x, y int) {
		live := 0
		for _, dir := range dirs {
			u, v := x + dir[0], y + dir[1]
			if u < 0 || v < 0 || u >= m || v >= n {
				continue
			}
			if board[u][v] == 1 || board[u][v] == -1 {
				live++
			}
		}
		if board[x][y] == 1 {
			if live < 2 || live > 3 {
				board[x][y] = -1
			}
		} else {
			if live == 3 {
				board[x][y] = 2
			}
		}
	}

	for i := range m {
		for j := range n {
			life(i, j)
		}
	}

	for i := range m {
		for j := range n {
			if board[i][j] == -1 {
				board[i][j] = 0 
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

// -1 is live to dead (that live)
// 2 is dead to live (dead)

// @lc code=end

