/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}

	var backtrack func(cur []int, start int)
	backtrack = func(cur []int, start int) {
		if len(cur) == k {
			res = append(res, append([]int{}, cur...))
		}

		for i := start; i < n+1; i++ {
			backtrack(append(cur, i), i+1)
		}
	}
	backtrack([]int{}, 1)
	return res
}

// @lc code=end

