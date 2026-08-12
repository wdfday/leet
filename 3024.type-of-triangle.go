/*
 * @lc app=leetcode id=3024 lang=golang
 *
 * [3024] Type of Triangle
 */

// @lc code=start
func triangleType(nums []int) string {
	slices.Sort(nums)
	if nums[0] + nums[1] <= nums[2] {
		return "none"
	} 
	if nums[0] == nums[2] {
	return "equilateral"
}

if nums[0] == nums[1] || nums[1] == nums[2] {
	return "isosceles"
}
	return "scalene"
}
// @lc code=end

