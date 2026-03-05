/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {

	n := len(nums)
	sort.Ints(nums)

	// Initialize result with the first triplet sum
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Update result if current sum is closer to target
			if abs(sum-target) < abs(result-target) {
				result = sum
			}

			// If exact match, return immediately
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

