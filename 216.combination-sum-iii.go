/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	var backtrack func(start int, path []int, sum int)
	backtrack = func(start int, path []int, sum int) {
		if len(path) == k && sum == n {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			backtrack(i+1, path, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtrack(1, []int{}, 0)

	return res

}

// @lc code=end

