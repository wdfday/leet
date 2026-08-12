/*
 * @lc app=leetcode id=1030 lang=golang
 *
 * [1030] Matrix Cells in Distance Order
 */

// @lc code=start
func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	res := make([][]int, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = append(res, []int{i, j})
		}
	}

	dist := func(cell []int) int {
		return abs(cell[0]-rCenter) + abs(cell[1]-cCenter)
	}

	slices.SortFunc(res, func(a, b []int) int {
		return dist(a) - dist(b)
	})

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
// @lc code=end

