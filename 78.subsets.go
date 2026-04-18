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
			cur = append(cur, nums[i])
			sub(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	sub(make([]int, 0), 0)

	return res
}

// @lc code=end

