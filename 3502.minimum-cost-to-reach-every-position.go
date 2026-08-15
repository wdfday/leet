/*
 * @lc app=leetcode id=3502 lang=golang
 *
 * [3502] Minimum Cost to Reach Every Position
 */

// @lc code=start
func minCosts(cost []int) []int {
	m := cost[0]

	for i := 1; i < len(cost); i++ {
		if cost[i] < m {
			m = cost[i]
		} else {
			cost[i] = m
		}
	}


    return cost
}
// @lc code=end

