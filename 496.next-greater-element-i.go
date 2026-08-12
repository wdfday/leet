/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			if stack[len(stack) - 1] > nums2[i] {
				m[nums2[i]] = stack[len(stack) - 1]
				break
			} else {
				stack = stack[:len(stack) - 1]
			}
		}
		
		if len(stack) == 0 {
			m[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])

	}

	for i, v := range nums1 {
		nums1[i] = m[v]
	}

    return nums1
}
// @lc code=end

