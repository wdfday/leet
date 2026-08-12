/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})

	if k > len(nums)-1 {
		k = len(nums) - 1
	}

	for i := 0; i <= k; i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}

	for i := k + 1; i < len(nums); i++ {
		delete(m, nums[i-k-1])

		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}

	}
	return false
}

// @lc code=end

