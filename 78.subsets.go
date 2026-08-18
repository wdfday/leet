/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}

	var sub func(cur []int, start int)
	sub = func(cur []int, start int) {
		res = append(res, append([]int{}, cur...))

		for i := start; i < len(nums); i++ {
			sub(append(cur, nums[i]), i+1)
		}
	}

	sub([]int{}, 0)

	return res
}

// @lc code=end

