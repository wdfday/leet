/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort theo start point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		lastIdx := len(result) - 1

		if curr[0] <= result[lastIdx][1] {
			result[lastIdx][1] = max(result[lastIdx][1], curr[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

