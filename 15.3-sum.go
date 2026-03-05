/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {

	result := [][]int{}

	n := len(nums)
	if n < 3 {
		return result
	}

	// Sort the array to use two-pointer technique
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate values for the first number
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // Skip duplicate values for the second number
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // Skip duplicate values for the third number
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result

}

// @lc code=end

