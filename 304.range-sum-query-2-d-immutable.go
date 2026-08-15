/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	prefix := make([][]int, m+1)
	for i := range m+1 {
		prefix[i] = make([]int, n+1)
	}
    
	for i := range m {
		for j := range n {
			prefix[i+1][j+1] =
				matrix[i][j] +
				prefix[i][j+1] +
				prefix[i+1][j] -
				prefix[i][j]
		}
	}

	return NumMatrix{matrix:prefix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1] + this.matrix[row1][col1] - this.matrix[row2+1][col1] - this.matrix[row1][col2+1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

