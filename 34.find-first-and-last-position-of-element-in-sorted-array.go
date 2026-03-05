/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {

	left := 0
	right := len(nums) - 1
	result := []int{-1, -1}

	// Find the first occurrence
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return result
	}
	result[0] = left

	// Find the last occurrence
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	result[1] = right

	return result
}

// @lc code=end

