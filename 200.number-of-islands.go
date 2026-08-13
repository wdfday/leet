/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])

	dirs := [][2]int{{1,0}, {0,1}, {-1, 0}, {0, -1}}

	var bfs func(x, y int)
	bfs = func(x, y int) {
		grid[x][y] = '0'
		for _, dir := range dirs {
			u, v := x + dir[0], y + dir[1]
			if u < 0 || v < 0 || u >= m || v >= n {
				continue
			} 
			if grid[u][v] == '0' {
				continue
			} 
			bfs(u, v)
		}
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				// bfs
				bfs(i, j)
				res ++
			}
		}
	}
    

	return res
}
// @lc code=end

