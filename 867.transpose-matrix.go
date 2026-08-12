/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 */

// @lc code=start
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			row[j] = matrix[j][i]
		}
		res = append(res, row)
	}
	return res
}

// @lc code=end

