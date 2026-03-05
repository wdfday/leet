/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}

	var backtrack func(start int, path []int, remaining int)
	backtrack = func(start int, path []int, remaining int) {
		if remaining == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, path, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{}, target)
	return result
}

// @lc code=end

