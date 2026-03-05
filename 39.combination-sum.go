/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sắp xếp để break sớm
	result := [][]int{}

	var backtrack func(start int, path []int, remaining int)
	backtrack = func(start int, path []int, remaining int) {
		if remaining == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				break // Dừng sớm vì đã sort
			}
			path = append(path, candidates[i])
			backtrack(i, path, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{}, target)
	return result

}

// @lc code=end

