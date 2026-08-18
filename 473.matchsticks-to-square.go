/*
 * @lc app=leetcode id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}

	if sum%4 != 0 {
		return false
	}

	edge := sum / 4

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	if matchsticks[0] > edge {
		return false
	}

	sides := [4]int{}

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return sides[0] == edge &&
				sides[1] == edge &&
				sides[2] == edge &&
				sides[3] == edge
		}

		v := matchsticks[idx]

		for i := 0; i < 4; i++ {
			if sides[i]+v > edge {
				continue
			}

			// symmetric states
			if i > 0 && sides[i] == sides[i-1] {
				continue
			}

			sides[i] += v

			if dfs(idx + 1) {
				return true
			}

			sides[i] -= v

			// All empty sides are equivalent.
			if sides[i] == 0 {
				break
			}
		}

		return false
	}

	return dfs(0)
}
// @lc code=end

