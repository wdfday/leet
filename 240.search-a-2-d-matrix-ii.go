/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := range m {
		if target > matrix[i][n-1] {
			continue
		} else if target < matrix[i][0] {
			break
		} else {
			l, r := 0, n - 1
			for l <= r {
				mid := l + (r - l)/2
				if matrix[i][mid] == target {
					return true
				} else if target > matrix[i][mid] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}
	}
	return false
}
// @lc code=end

func searchMatrix(matrix [][]int, target int) bool {
    ROWS, COLS := len(matrix), len(matrix[0])
    r, c := 0, COLS-1
    for r < ROWS && c >= 0 {
        if matrix[r][c] == target {
            return true
        }
        if matrix[r][c] > target {
            c--
        } else {
            r++
        }
    } 
    return false
}