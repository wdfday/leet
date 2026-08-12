/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}

	var res []int

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}

	return res

}

// @lc code=end

