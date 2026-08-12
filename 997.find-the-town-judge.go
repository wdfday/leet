/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	score := make([]int, n+1)
	for _, t := range trust {
		score[t[0]]--
		score[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if score[i] == n-1 {
			return i
		}
	}
	return -1
}

// @lc code=end

