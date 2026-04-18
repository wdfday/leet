/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {

	var res [][]int
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			res = append(res, perm)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return res

}

// @lc code=end

