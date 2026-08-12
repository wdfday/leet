/*
 * @lc app=leetcode id=3731 lang=golang
 *
 * [3731] Find Missing Elements
 */

// @lc code=start
func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1]+1 {
			for k := nums[i-1] + 1; k < nums[i]; k++ {
				res = append(res, k)
			}
		}
	}
	return res
}

// @lc code=end

