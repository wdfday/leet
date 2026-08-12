/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i+k >= 0 && i+k < m && j+l >= 0 && j+l < n {
						sum += img[i+k][j+l]
						count++
					}
				}
			}
			res[i][j] = sum / count
		}
	}

	return res

}

// @lc code=end

