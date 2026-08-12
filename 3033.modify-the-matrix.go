/*
 * @lc app=leetcode id=3033 lang=golang
 *
 * [3033] Modify the Matrix
 */

// @lc code=start
func modifiedMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	for j := 0; j < n; j ++ {
		k := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] > k {
				k =matrix[i][j]
			}
		}
		for i := 0; i < m; i++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = k
			}
		}
		
	}
    return matrix
}
// @lc code=end

