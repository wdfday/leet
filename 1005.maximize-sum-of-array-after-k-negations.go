/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)
	res := 0
	minabs := abs(nums[0])

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			res -= nums[i]
			k--
		} else {
			res += nums[i]
		}
		if v := abs(nums[i]); v < minabs {
			minabs = v
		}
	}

	if k%2 == 1 {
		res -= 2 * minabs
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
// @lc code=end

