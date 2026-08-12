/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}

	var res []int

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] {
			res = append(res, nums2[i])
			m[nums2[i]] = false
		}
	}

	return res

}

// @lc code=end

