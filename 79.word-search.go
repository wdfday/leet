/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(idx, i, j int) bool
	dfs = func(idx, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = '#' // đánh dấu đang dùng

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		found := false
		for _, d := range dirs {
			if dfs(idx+1, i+d[0], j+d[1]) {
				found = true
				break
			}
		}

		board[i][j] = tmp // backtrack, trả lại giá trị gốc
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

