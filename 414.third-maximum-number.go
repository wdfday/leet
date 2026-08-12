/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {

	var max1, max2, max3 *int

	for _, num := range nums {
		if max1 != nil && *max1 == num || max2 != nil && *max2 == num || max3 != nil && *max3 == num {
			continue
		}

		if max1 == nil || num > *max1 {
			max3 = max2
			max2 = max1
			max1 = &num
		} else if max2 == nil || num > *max2 {
			max3 = max2
			max2 = &num
		} else if max3 == nil || num > *max3 {
			max3 = &num
		}
	}

	if max3 != nil {
		return *max3
	}

	return *max1
}

// @lc code=end

