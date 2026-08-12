/*
 * @lc app=leetcode id=3010 lang=golang
 *
 * [3010] Divide an Array Into Subarrays With Minimum Cost I
 */

// @lc code=start
func minimumCost(nums []int) int {
	x1, x2 := nums[1], nums[2]

	for i := 3; i < len(nums); i++ {
		if nums[i] < max(x1,x2) {
			x1, x2 = nums[i], min(x1,x2)
		}
	}


    return nums[0] + x1 + x2
}
// @lc code=end

