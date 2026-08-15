/*
 * @lc app=leetcode id=3507 lang=golang
 *
 * [3507] Minimum Pair Removal to Sort Array I
 */

// @lc code=start
func minimumPairRemoval(nums []int) int {
	ops := 0
	for !isSorted(nums) {
		nums = mergeMinPair(nums)
		ops++
	}
	return ops
}

func mergeMinPair(nums []int) []int {
	minIdx, minSum := 0, nums[0]+nums[1]
	for i := 1; i < len(nums)-1; i++ {
		if s := nums[i] + nums[i+1]; s < minSum {
			minIdx, minSum = i, s
		}
	}

	merged := make([]int, 0, len(nums)-1)
	merged = append(merged, nums[:minIdx]...)
	merged = append(merged, minSum)
	merged = append(merged, nums[minIdx+2:]...)
	return merged
}

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
// @lc code=end

