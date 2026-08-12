/*
 * @lc app=leetcode id=3789 lang=golang
 *
 * [3789] Minimum Cost to Acquire Required Items
 */

// @lc code=start
func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	var res int64
	m, n := max(need1, need2), min(need1, need2)
	res = int64(need1*cost1 + need2*cost2)

	res1 := int64(m * costBoth)
	res2 := int64(n*costBoth + cost1*(need1-n) + cost2*(need2-n))

	return min(res, res1, res2)
}

// @lc code=end

