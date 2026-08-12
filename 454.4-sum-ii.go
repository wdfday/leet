/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)

	// build sum of first two arrays
	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}

	// check against last two arrays
	res := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			res += m[-(c + d)]
		}
	}

	return res
}

// @lc code=end

