/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			matrix[first][i], matrix[last-offset][first], matrix[last][last-offset], matrix[i][last] =
				matrix[last-offset][first], matrix[last][last-offset], matrix[i][last], matrix[first][i]
		}
	}
}

// @lc code=end

// func rotate(matrix [][]int) {
// 	n := len(matrix)

// 	// Transpose the matrix
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
// 		}
// 	}

// 	// Reverse each row
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n/2; j++ {
// 			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
// 		}
// 	}

// }