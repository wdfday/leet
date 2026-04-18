/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	t, d := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1

	if matrix[t][0] > target || matrix[d][r] < target {
		return false
	}

	for t < d {
		mid := (t + d + 1) / 2
		if matrix[mid][0] > target {
			d = mid - 1
		} else {
			t = mid
		}
	}

	for l < r {
		mid := (l + r + 1) / 2
		if matrix[t][mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return matrix[t][l] == target

}

// @lc code=end

