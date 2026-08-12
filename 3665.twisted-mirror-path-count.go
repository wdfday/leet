/*
 * @lc app=leetcode id=3665 lang=golang
 *
 * [3665] Twisted Mirror Path Count
 */

// @lc code=start
const MOD = 1_000_000_007

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	var solve func(int, int, bool) int
	solve = func(i, j int, isR bool) int {
		if isR {
			j++
		} else {
			i++
		}

		if i >= m || j >= n {
			return 0
		}
		if grid[i][j] == 0 {
			return res[i][j]
		}
		return solve(i, j, !isR)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				res[i][j] = 1
				continue
			}

			res[i][j] = (solve(i, j, true) + solve(i, j, false)) % MOD
		}
	}
	return res[0][0]
}

// @lc code=end

