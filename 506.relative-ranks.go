/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
func findRelativeRanks(score []int) []string {
	op := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	indices := make([]int, len(score))
	for i := 0; i < len(indices); i++ {
		indices[i] = i
	}
	res := make([]string, len(score))
	sort.Slice(indices, func(i, j int) bool {
		return score[indices[i]] > score[indices[j]]
	})

	for i := 0; i < len(indices); i++ {
		if i == 0 {
			res[indices[i]] = op[0]
		} else if i == 1 {
			res[indices[i]] = op[1]
		} else if i == 2 {
			res[indices[i]] = op[2]
		} else {
			res[indices[i]] = fmt.Sprintf("%v", i+1)
		}
	}

	return res
}

// @lc code=end

