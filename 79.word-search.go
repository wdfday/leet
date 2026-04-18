/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(idx, i, j, prevI, prevJ int) bool
	dfs = func(idx, i, j, prevI, prevJ int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || (i == prevI && j == prevJ) {
			return false
		}
		if board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if dfs(idx+1, ni, nj, i, j) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j, -1, -1) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

