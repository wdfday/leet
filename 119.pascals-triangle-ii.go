/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

// @lc code=end

