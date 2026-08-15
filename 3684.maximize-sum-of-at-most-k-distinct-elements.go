/*
 * @lc app=leetcode id=3684 lang=golang
 *
 * [3684] Maximize Sum of At Most K Distinct Elements
 */

// @lc code=start
func maxKDistinct(nums []int, k int) []int {
	slices.Sort(nums)
	res := []int{0}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != res[len(res)-1] {
			if k == 0 {
				break
			}
			res = append(res, nums[i])
			k--
		}
	}

	return res[1:]
    
}
// @lc code=end

