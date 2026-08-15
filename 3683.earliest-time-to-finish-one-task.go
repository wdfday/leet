/*
 * @lc app=leetcode id=3683 lang=golang
 *
 * [3683] Earliest Time to Finish One Task
 */

// @lc code=start
func earliestTime(tasks [][]int) int {
	res := math.MaxInt
	for _, task := range tasks {
		res = min(task[0]+task[1], res)
	}
	return res
}
// @lc code=end

