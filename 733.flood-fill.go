/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m, n := len(image), len(image[0])
	pref := image[sr][sc]

	dirs := [][2]int{{0,1}, {1,0}, {0,-1}, {-1, 0}}

	var fill func(x, y int) 
	fill = func(x, y int) {
		image[x][y] = color

		for _, dir := range dirs {
			u, v := x + dir[0], y + dir[1]
			if u < 0 || v < 0 || u >= m || v >= n {
				continue
			}
			if image[u][v] != pref || image[u][v] == color  {
				continue
			}
			fill(u, v)
		}
	}

	fill(sr, sc)
	return image
}
// @lc code=end

