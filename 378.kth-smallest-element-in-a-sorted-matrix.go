/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])

	lo := matrix[0][0]
	hi := matrix[m-1][n-1]

	countLE := func(x int) int {
		i, j := m-1, 0
		cnt := 0

		for i >= 0 && j < n {
			if matrix[i][j] <= x {
				cnt += i + 1
				j++
			} else {
				i--
			}
		}

		return cnt
	}

	for lo < hi {
		mid := lo + (hi-lo)/2

		if countLE(mid) < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
// @lc code=end

